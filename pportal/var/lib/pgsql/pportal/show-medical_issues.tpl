<h1 class="titlebg">Medical Issues</h1>
<table>
  <tr><th>Title</th></tr>
  <%= 
(let loop ((res data) (html ""))
  (cond
    ((null? res) html)
    (else
      (loop (cdr res)
        (format #f "~a~%<tr><td>~a</td></tr>"
          html
          (cdr (car (cdr (car res)))))))))
%>
</table>
