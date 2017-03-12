<!DOCTYPE html>
<html>
<head>
	<title></title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css">
</head>
<h1>Register Page</h1>
{{.}}
<form method="post" action="/register">
    <label for="name">User name</label>
    <input type="text" id="name" name="name" value="nabil">
    <label for="password">Password</label>
    <input type="password" id="password" name="password" value="nabil">
    <label for="email">Email</label>
    <input type="email" id="email" name="email" value="nabil">
    <button type="submit">Register</button>
    
</form>