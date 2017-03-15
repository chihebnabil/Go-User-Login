<!DOCTYPE html>
<html>
<head>
    <title></title>
<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css">
</head>
<div class="row">
   <div class="container">
       <h1>Edit Profile Information Page</h1>
<form method="POST" action="/edit">
    <label for="full_name">Full Name</label>
    <input class="form-control" required type="text" id="full_name" name="full_name" value="{{index . 1}}">
    <br>
    <label for="address">Address</label>
    <input class="form-control" required type="text" id="autocomplete" name="address"  value="{{index . 3}}">
    <br>
    <label for="phone">Phone</label>
    <input class="form-control" required type="text" id="phone" name="phone" value="{{index . 2}}" required>
    <br>
    <button type="submit" class="btn btn-primary">Update</button>
</form>
   </div>    
</div>
<script type="text/javascript">

function initAutocomplete() {
  autocomplete = new google.maps.places.Autocomplete(
     /** @type {!HTMLInputElement} */(document.getElementById('autocomplete')),
     {types: ['geocode']});
}

</script>
<script src="https://maps.googleapis.com/maps/api/js?key=AIzaSyD9xKUDq37EGcZtPz-ZxmyQo5bsLi2EyYc&libraries=places&callback=initAutocomplete"
       async defer></script>
</body>