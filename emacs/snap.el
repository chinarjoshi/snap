;;; snap.el --- Semantic Note Auto-Parsers -*- lexical-binding: t; -*-

(require 'snap-core nil t)  ;; loaded by use-package :config if missing

(defun snap-workout-result-exercises (result)
  "Extract exercises from a workout parse result."
  (cdr (assq 'exercises result)))

(defun snap-workout-result-prose (result)
  "Extract prose lines from a workout parse result."
  (cdr (assq 'prose result)))

(provide 'snap)
;;; snap.el ends here
