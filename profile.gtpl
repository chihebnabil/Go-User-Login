<!DOCTYPE html>
<html>
<head>
    <title></title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css">
</head>
<div class="row">
   <div class="container">
       <h1>Fill Profile Information Page</h1>
<form method="POST" action="/profil">
    <label for="full_name">Full Name</label>
    <input class="form-control" required type="text" id="full_name" name="full_name" value="nabil">
    <br>
    <label for="address">Address</label>
    <input class="form-control" required type="text" id="address" name="address" value="nabil">
    <br>
    <label for="phone">Phone</label>
    <input class="form-control" required type="text" id="phone" name="phone" value="" required>
    <br>
    <button type="submit" class="btn btn-primary">Update</button>
</form>
   </div>    
</div>
