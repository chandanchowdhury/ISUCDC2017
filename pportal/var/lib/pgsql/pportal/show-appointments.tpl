<h1 class="titlebg">Appointments</h1>
<table>
  <tr>
    <th>Title</th>
    <th>Date</th>
    <th>Start Time</th>
    <th>End Time</th>
    <th>Comments</th>
  </tr>
  <%= 
(let loop ((res data) (html ""))
  (cond
    ((null? res) html)
    (else
      (loop (cdr res)
        (format #f 
          "~a~%<tr><td>~a</td><td>~a</td><td>~a</td><td>~a</td><td>~a</td></tr>"
          html
          (cdr (car (cdr (car res))))
          (cdr (car (cdr (cdr (car res)))))
          (cdr (car (cdr (cdr (cdr (cdr (car res)))))))
          (cdr (car (cdr (cdr (cdr (cdr (cdr (car res))))))))
          (cdr (car (cdr (cdr (cdr (car res)))))))))))
%>
</table>
