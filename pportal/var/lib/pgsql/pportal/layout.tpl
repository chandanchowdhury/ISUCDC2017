<html>
	<head>
		<title>Patient portal</title>
	<style>
		table,td,th { border: 1px solid black; background-color: tan;}
		table { width: 100% ;}
		body { background-color: grey; padding-left: 10%; padding-right: 10%}
		.titlebg {background-color: tan; padding-left: 3%; border: 1px solid black;}
		ul.nav,li.nav { list-style: none; display: inline; padding-right: 20px;}
	</style>
	</head>
	<body>
		<div class="navbar">
			<ul class="nav">
				<li class="nav">
					<a href="/">Home</a>
				</li>
				<li class="nav">
					<a href="/showdata/allergies">Allergies</a>
				</li>
				<li class="nav">
					<a href="/showdata/appointments">Appointments</a>
				</li>
				<li class="nav">
					<a href="/showdata/medical_issues">Medical Issues</a>
				</li>
				<li class="nav">
					<a href="/showdata/medications">Medications</a>
				</li>
				<li class="nav">
					<a href="/login">Login</a>
				</li>
				<li class="nav">
					<a href="/logout">Logout</a>
				</li>
			</ul>
		</div>
		<%= page %>
	</body>
</html>
