(use-modules (dbi dbi))

(define db-query
  (lambda (query)
    (let ((conn (dbi-open "postgresql" "postgres::postgres:tcp:localhost:5432")))
      (dbi-query conn query)
      (let get-row ((rows '()) (row (dbi-get_row conn)))
        (cond
          (row
            (get-row (append rows (cons row '())) (dbi-get_row conn)))
          (else
	    ;(dbi-close conn)
            (format #t "Query: ~a; Res: ~a~%" query rows)
            rows))))))
