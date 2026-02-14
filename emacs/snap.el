;;; snap.el --- Structured Note Auto-Parsers -*- lexical-binding: t; -*-

(require 'snap-core)  ;; loads the dynamic module

(defun snap-workout-result-exercises (result)
  "Extract exercises from a workout parse result."
  (cdr (assq 'exercises result)))

(defun snap-workout-result-prose (result)
  "Extract prose lines from a workout parse result."
  (cdr (assq 'prose result)))

(provide 'snap)
;;; snap.el ends here
