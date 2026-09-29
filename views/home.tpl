{{template "partials/header.tpl" .}}


<h2>
    Search Events
</h2>


<form action="/events" method="GET">


<input 
    id="cityInput"
    name="city"
    placeholder="Search city"
    autocomplete="off"
/>


<div id="suggestions"></div>


<input 
    type="hidden"
    id="placeId"
    name="placeId"
/>


<input
    type="hidden"
    id="countryCode"
    name="countryCode"
/>


<button type="submit">
    Search
</button>


</form>



<script src="/static/js/app.js"></script>


{{template "partials/footer.tpl" .}}