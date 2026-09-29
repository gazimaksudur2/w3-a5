let sessionToken = crypto.randomUUID();

const input = document.getElementById("cityInput");

const suggestionBox = document.getElementById("suggestions");

input.addEventListener("input", async function () {
  let value = input.value.trim();

  if (value.length < 2) {
    suggestionBox.innerHTML = "";
    return;
  }

  let response = await fetch(
    "/api/locations/autocomplete?input=" +
      value +
      "&sessionToken=" +
      sessionToken,
  );

  let data = await response.json();

  suggestionBox.innerHTML = "";

  data.forEach((item) => {
    let div = document.createElement("div");

    div.innerText = item.description;

    div.onclick = function () {
      input.value = item.description;

      document.getElementById("placeId").value = item.placeId;

      suggestionBox.innerHTML = "";
    };

    suggestionBox.appendChild(div);
  });
});
