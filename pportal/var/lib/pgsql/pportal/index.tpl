<h1>ISEmr Patient Portal</h1>
</p>
<%= (cond
      ((not (string-null? lname))
        (format #f "Welcome, ~a ~a. Use the links above to navigate." fname lname))
      (else
        "Please log in to access the patient portal.")))
%>
</p>
