;;; snap.el --- Semantic Note Auto-Parsers -*- lexical-binding: t; -*-

;; Auto-build snap-core dynamic module if not already loaded.
(unless (require 'snap-core nil t)
  (let ((default-directory (file-name-directory (locate-library "snap"))))
    (shell-command "cd .. && git submodule update --init --recursive && cd emacs && make")
    (require 'snap-core)))

(require 'cl-lib)

(defun snap--get (key alist)
  "Get KEY from ALIST, coercing vectors to lists."
  (append (cdr (assq key alist)) nil))

;;; Formatting

(defface snap-table
  '((t :foreground "#8a8a8a" :family "Monospace"))
  "Face for rendered SNAP workout tables.")

(defun snap--format-row (ex max-sets)
  "Format exercise EX as a table row, padded to MAX-SETS columns."
  (let ((prev-weight nil)
        (sets (snap--get 'sets ex)))
    (append
     (list (cdr (assq 'name ex)))
     (cl-loop for s in sets
              for reps = (cdr (assq 'reps s))
              for weight = (cdr (assq 'weight s))
              collect (cond
                       ((= weight 0) (number-to-string reps))
                       ((equal weight prev-weight) (number-to-string reps))
                       (t (format "%dx%d" reps weight)))
              do (setq prev-weight weight))
     (make-list (- max-sets (length sets)) ""))))

(defun snap--format-table (exercises)
  "Format EXERCISES as an aligned org table string with notes."
  (let* ((max-sets (cl-reduce #'max (mapcar (lambda (ex) (length (snap--get 'sets ex)))
                                            exercises)))
         (rows (mapcar (lambda (ex) (snap--format-row ex max-sets)) exercises))
         (header (cons "" (mapcar #'number-to-string (number-sequence 1 max-sets))))
         (raw (concat
               "|" (mapconcat #'identity header "|") "|\n"
               "|-\n"
               (mapconcat (lambda (row) (concat "|" (mapconcat #'identity row "|") "|"))
                          rows "\n")))
         (table (with-temp-buffer
                  (org-mode)
                  (insert raw)
                  (org-table-align)
                  (string-trim (buffer-string))))
         (notes (cl-loop for ex in exercises
                         for name = (cdr (assq 'name ex))
                         for ex-notes = (cl-loop for s in (snap--get 'sets ex)
                                                 for note = (cdr (assq 'note s))
                                                 when (and note (> (length note) 0))
                                                 collect note)
                         when ex-notes
                         collect (format "- %s :: %s" name (mapconcat #'identity ex-notes ". ")))))
    (concat
     (propertize table 'face 'snap-table)
     (when notes
       (concat "\n" (mapconcat #'identity notes "\n"))))))

;;; Overlays

(defun snap--render-overlay (beg end exercises)
  "Put an overlay from BEG to END showing the workout table for EXERCISES."
  (let ((ov (make-overlay beg end)))
    (overlay-put ov 'snap-workout t)
    (overlay-put ov 'display (snap--format-table exercises))
    ov))

(defun snap--overlay-near-point ()
  "Find a snap overlay near point (on previous line or adjacent)."
  (cl-find-if (lambda (ov) (overlay-get ov 'snap-workout))
              (overlays-in (max (- (point) 2) (point-min)) (point))))

(defun snap--remove-overlays ()
  "Remove all snap overlays in the current buffer."
  (dolist (ov (overlays-in (point-min) (point-max)))
    (when (overlay-get ov 'snap-workout)
      (delete-overlay ov))))

;;; Drawer operations

(defun snap--paragraph-before-point ()
  "Return (beg end text) of paragraph before point."
  (save-excursion
    (skip-chars-backward "\n")
    (let* ((end (point))
           (start (if (re-search-backward "\n\n" nil t) (match-end 0) (point-min))))
      (list start end (string-trim (buffer-substring-no-properties start end))))))

(defun snap--render-workout-at (beg end text)
  "Wrap TEXT from BEG to END in a :SNAP: drawer and overlay.
Prose lines are kept as regular text above the drawer."
  (let* ((result (snap-workout text))
         (exercises (snap--get 'exercises result)))
    (when exercises
      (let* ((prose-set (mapcar #'downcase (snap--get 'prose result)))
             (lines (split-string text "\n"))
             workout-lines prose-lines)
        (dolist (line lines)
          (if (member (downcase (string-trim line)) prose-set)
              (push line prose-lines)
            (push line workout-lines)))
        (setq workout-lines (nreverse workout-lines)
              prose-lines (nreverse prose-lines))
        (delete-region beg end)
        (goto-char beg)
        (when prose-lines
          (insert (mapconcat #'identity prose-lines "\n") "\n"))
        (let ((snap-beg (point)))
          (insert ":SNAP:\n" (mapconcat #'identity workout-lines "\n") "\n:END:")
          (snap--render-overlay snap-beg (point) exercises)
          (insert "\n")))
      t)))

(defun snap--derender-at-point ()
  "If there is a :SNAP: overlay near point, derender it."
  (let ((ov (snap--overlay-near-point)))
    (when ov
      (let* ((beg (overlay-start ov))
             (end (overlay-end ov))
             (text (buffer-substring-no-properties beg end)))
        (delete-overlay ov)
        (when (string-match "\\`:SNAP:\n\\(\\(?:.\\|\n\\)*?\\)\n:END:\\'" text)
          (let ((raw (match-string 1 text)))
            (delete-region beg (min (1+ end) (point-max)))
            (goto-char beg)
            (insert raw))))
      t)))

(defun snap--render-buffer ()
  "Find all :SNAP: drawers and render overlays."
  (save-excursion
    (goto-char (point-min))
    (while (re-search-forward "^:SNAP:\n" nil t)
      (let ((beg (match-beginning 0)))
        (when (re-search-forward "^:END:" nil t)
          (let* ((end (match-end 0))
                 (trimmed (string-trim (buffer-substring-no-properties (+ beg 6) (- end 5))))
                 (exercises (snap--get 'exercises (snap-workout trimmed))))
            (when exercises
              (snap--render-overlay beg end exercises))))))))

;;; Hooks

(defun snap--on-double-enter ()
  "Run SNAP parsers on the preceding paragraph after double-enter."
  (when (looking-back "\n\n" (max 1 (- (point) 2)))
    (let* ((full-end (point))
           (para (snap--paragraph-before-point))
           (beg (nth 0 para))
           (text (nth 2 para)))
      (when (and (> (length text) 0)
                 (not (string-match-p ":SNAP:" text)))
        (snap--render-workout-at beg full-end text)))))

(defun snap--maybe-derender ()
  "In pre-command-hook: if about to backspace near a snap overlay, derender."
  (when (and (memq this-command '(delete-backward-char
                                  org-delete-backward-char
                                  backward-delete-char-untabify))
             (snap--derender-at-point))
    (setq this-command 'ignore)))

;;; Minor mode

;;;###autoload
(define-minor-mode snap-mode
  "Minor mode for rendering SNAP drawers as formatted tables."
  :lighter " snap"
  (if snap-mode
      (progn
        (snap--render-buffer)
        (add-hook 'post-command-hook #'snap--on-double-enter nil t)
        (add-hook 'pre-command-hook #'snap--maybe-derender nil t))
    (snap--remove-overlays)
    (remove-hook 'post-command-hook #'snap--on-double-enter t)
    (remove-hook 'pre-command-hook #'snap--maybe-derender t)))

(provide 'snap)
;;; snap.el ends here
