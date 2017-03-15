<!DOCTYPE html>
<html>
<head>
    <title></title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css">
</head>
<h1>My Profile</h1>
<hr>
<form method="post" action="/logout">
    <button class="btn btn-default" type="submit">Logout</button>
</form>
<div class="row">
	<div class="container">
<p>User: {{index . 0}}</p>
	  <p>FullName :{{index . 1}}</p>
	  <p>Phone : {{index . 2}}</p>
	  <p>Address : {{index . 3}}</p>	
	   <p>Created : {{index . 4}}</p>	
	  <a class="btn btn-default" href="/edit">Edit</a>
	</div>
</div>