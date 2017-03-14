<!DOCTYPE html>
<html>
<head>
	<title></title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css">
</head>
<body>
<div class="row">
   <div class="container">
<h1>Login</h1>
<form method="post" action="/login">
    <label for="email">User name</label>
    <input class="form-control" type="email" id="email" name="email" value="chiheb.design@gmail.com">
    <label for="password">Password</label>
    <input class="form-control" type="password" id="password" name="password" value="nabil">
    <button class="btn btn-primary" type="submit">Login</button>
    <a class="btn btn-default" href="/register">Register</a>
    <a class="btn btn-default" href="/lost">Lost Password</a>
</form>
</div>
</div>
</body>
</html>