let sessionToken = crypto.randomUUID();


const input = document.getElementById("cityInput");

const suggestionBox =
document.getElementById("suggestions");



input.addEventListener(
"input",
async function(){


    let value = input.value.trim();



    if(value.length < 2){

        suggestionBox.innerHTML="";

        return;

    }



    suggestionBox.innerHTML =
    `
    <div class="p-3 text-gray-500">
        Searching...
    </div>
    `;



    let response =
    await fetch(
        "/api/locations/autocomplete?input="
        +
        encodeURIComponent(value)
        +
        "&sessionToken="
        +
        sessionToken
    );



    let data =
    await response.json();



    suggestionBox.innerHTML="";



    data.forEach(item=>{


        let div =
        document.createElement("div");



        div.className =
        "p-3 cursor-pointer hover:bg-gray-100";



        div.innerText =
        item.description;



        div.onclick =
        async function(){



            input.value =
            item.description;



            suggestionBox.innerHTML="";



            let locationResponse =
            await fetch(
                "/api/locations/"
                +
                item.placeId
                +
                "?sessionToken="
                +
                sessionToken
            );



            let location =
            await locationResponse.json();



            document.getElementById(
                "countryCode"
            ).value =
            location.countryCode;



            input.value =
            location.city;


        };



        suggestionBox.appendChild(div);


    });


});