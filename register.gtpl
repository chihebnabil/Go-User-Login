<!DOCTYPE html>
<html>
<head>
	<title></title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css">
</head>
<div class="row">
   <div class="container">
<h1>Register Page</h1>
{{.}}
<form method="post" action="/register">
    <label for="name">User name</label>
    <input class="form-control" type="text" id="name" name="name" value="nabil">
    <label for="password">Password</label>
    <input class="form-control" type="password" id="password" name="password" value="nabil">
    <label for="email">Email</label>
    <input class="form-control" type="email" id="email" name="email" value="chiheb.design@gmail.com">
    <button type="submit">Register</button>
    
</form>
</div>
</div>