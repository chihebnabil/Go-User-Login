<!DOCTYPE html>
<html>
<head>
	<title></title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css">
</head>
<body>
<div class="row">
   <div class="container">
<h1>Lost Password</h1>
{{.}}
<form method="post" action="/lost">
       <label for="email">Email</label>
    <input class="form-control" type="email" id="email" name="email" value="">
    <button type="submit">Send</button>
   
</form>

</div>
</div>
</body>
</html>