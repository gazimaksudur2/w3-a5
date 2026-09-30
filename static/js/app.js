let sessionToken = crypto.randomUUID();


const input =
    document.getElementById("cityInput");


const suggestionBox =
    document.getElementById("suggestions");


const countryCodeInput =
    document.getElementById("countryCode");



let debounceTimer;



// Search cities

input.addEventListener(
    "input",
    function() {


        clearTimeout(debounceTimer);



        let value =
            input.value.trim();



        if (value.length < 2) {

            suggestionBox.innerHTML = "";

            return;

        }



        debounceTimer =
            setTimeout(
                () => searchCities(value),
                400
            );


    });




async function searchCities(value) {


    suggestionBox.innerHTML =
        `
    <div class="
        p-4
        text-slate-500
        text-center
    ">
        Searching cities...
    </div>
    `;



    try {


        let response =
            await fetch(
                "/api/locations/autocomplete?input=" +
                encodeURIComponent(value) +
                "&sessionToken=" +
                sessionToken
            );



        let data =
            await response.json();



        suggestionBox.innerHTML = "";



        if (!Array.isArray(data) || data.length === 0) {


            suggestionBox.innerHTML =
                `
            <div class="
                p-4
                text-slate-500
                text-center
            ">
                No cities found
            </div>
            `;


            return;

        }




        data.forEach(item => {


            let option =
                document.createElement("div");



            option.className =
                `
            px-5
            py-4
            cursor-pointer
            hover:bg-blue-50
            transition
            border-b
            border-slate-100
            `;



            option.innerHTML =
                `
            <div class="font-semibold text-slate-800">
                ${item.description}
            </div>

            <div class="text-xs text-slate-500 mt-1">
                Select this location
            </div>
            `;



            option.onclick =
                async function() {


                    input.value =
                        item.description;



                    suggestionBox.innerHTML =
                        `
                <div class="
                    p-4
                    text-slate-500
                    text-center
                ">
                    Loading location...
                </div>
                `;



                    try {


                        let locationResponse =
                            await fetch(
                                "/api/locations/" +
                                item.placeId +
                                "?sessionToken=" +
                                sessionToken
                            );



                        let location =
                            await locationResponse.json();



                        if (location.error) {

                            throw new Error(
                                location.error
                            );

                        }



                        input.value =
                            location.city;



                        countryCodeInput.value =
                            location.countryCode;



                        suggestionBox.innerHTML =
                            "";



                    } catch (error) {


                        suggestionBox.innerHTML =
                            `
                    <div class="
                        p-4
                        text-red-600
                    ">
                        Unable to load location
                    </div>
                    `;


                        console.error(error);

                    }



                };



            suggestionBox.appendChild(option);



        });



    } catch (error) {


        suggestionBox.innerHTML =
            `
        <div class="
            p-4
            text-red-600
            text-center
        ">
            Search failed
        </div>
        `;


        console.error(error);


    }



}




// Close dropdown when clicking outside


document.addEventListener(
    "click",
    function(event) {


        if (
            !input.contains(event.target) &&
            !suggestionBox.contains(event.target)
        ) {

            suggestionBox.innerHTML = "";

        }


    });