(define session-get-sid
  (lambda (sid field)
    (let ((res 
            (db-query 
              (format #f "SELECT ~a FROM sessions WHERE sid='~a'" field sid))))
      (cond
        ((null? res) "")
        (else (cdr (car (car res))))))))

(define session-set-sid
  (lambda (sid field value)
    (if (null? (db-query 
                        (format 
                          #f "SELECT ~a FROM sessions WHERE sid='~a'"
                          field sid)))
      (db-query (format 
		  #f 
		  "INSERT INTO sessions (sid,~a) VALUES ('~a','~a')"
                  field sid value)))
    (let ((res 
            (db-query 
              (format #f "UPDATE sessions SET ~a='~a' WHERE sid='~a'"
                      field value sid))))
      (cond
        ((null? res) res)
        (else (cdr (car (car res))))))))

(define session-get
  (lambda (rc field)
    (session-get-sid (session-sid rc) field)))

(define session-set
  (lambda (rc field value)
    (session-set-sid (session-sid rc) field value)))

(define session-sid
  (lambda (rc)
    (let ((s (:session rc 'check-and-spawn)))
      (cond
        ((string? s) (format #t "session-sid: ~a~%" s) s)
        (else "")))))
