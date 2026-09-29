{{template "partials/header.tpl" .}}


<section class="text-center py-20">


<h2 class="text-5xl font-bold tracking-tight">

Discover Events Around You

</h2>


<p class="mt-5 text-lg text-gray-500">

Find concerts, sports and experiences in your city.

</p>



<form 
action="/events"
method="GET"
class="mt-10 max-w-xl mx-auto bg-white p-6 rounded-3xl shadow-xl"
>


<input
id="cityInput"
name="city"
placeholder="Search city..."
class="
w-full
px-5
py-4
rounded-xl
border
focus:ring-2
focus:ring-blue-500
outline-none
"
/>



<div id="suggestions"
class="text-left mt-2">
</div>



<input
type="hidden"
name="countryCode"
id="countryCode"
/>



<button

class="
mt-5
w-full
bg-blue-600
hover:bg-blue-700
text-white
font-semibold
py-4
rounded-xl
transition
"

>

Search Events

</button>


</form>


</section>



<script src="/static/js/app.js"></script>


{{template "partials/footer.tpl" .}}