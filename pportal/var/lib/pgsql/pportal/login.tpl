<% (if (not cor) %>
<h2 style="color: red;">Incorrect login info.</h2>
<% ) %>
<div style="margin-top: 20px; padding-left: 10%">
	<form action="/do_login" method="GET">
		<label>Username</label>
		<input type="text" name="uname"/>
		<label>PIN</label>
		<input type="password" name="pin"/>
		<input type="Submit" value="Log in"/>
	</form>
</div>
