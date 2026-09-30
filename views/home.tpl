{{template "partials/header.tpl" .}}

<section class="relative overflow-visible px-6 sm:px-10 lg:px-16">
    <!-- Background decoration -->
    <div class="absolute inset-0 bg-gradient-to-br from-blue-50 via-white to-indigo-50 -z-10"></div>

    <div class="relative z-10 py-20 text-center">
        <div
            class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-blue-100 text-blue-700 text-sm font-semibold mb-8"
        >
            🌎 Explore the world of events
        </div>

        <h2 class="text-5xl md:text-7xl font-extrabold tracking-tight text-slate-900 leading-tight">
            Discover Events

            <br />

            <span class="bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent"> Around You </span>
        </h2>

        <p class="mt-6 max-w-2xl mx-auto text-lg text-slate-600 leading-relaxed">
            Find concerts, sports events, and unforgettable experiences in cities around the world.
        </p>

        <!-- Search Box -->

        <form
            action="/events"
            method="GET"
            class="relative z-[100] mt-12 max-w-3xl mx-auto bg-white rounded-3xl shadow-2xl border border-slate-100 p-3 flex flex-col md:flex-row gap-3"
        >
            <div class="flex-1 relative">
                <input
                    id="cityInput"
                    name="city"
                    placeholder="Search a city..."
                    autocomplete="off"
                    class="w-full px-6 py-5 rounded-2xl bg-slate-50 border border-slate-200 text-lg outline-none focus:ring-4 focus:ring-blue-100 focus:border-blue-500 transition"
                />

                <div
                    id="suggestions"
                    class="absolute left-0 right-0 top-full mt-3 bg-white rounded-2xl shadow-2xl border border-slate-200 overflow-hidden text-left z-[99999] max-h-80 overflow-y-auto"
                ></div>
            </div>

            <input type="hidden" name="countryCode" id="countryCode" />

            <button
                class="px-8 py-5 rounded-2xl bg-gradient-to-r from-blue-600 to-indigo-600 text-white font-bold text-lg shadow-lg hover:shadow-xl hover:scale-105 transition duration-300"
            >
                Search Events →
            </button>
        </form>
    </div>
</section>

<!-- Categories -->

<section class="mt-16 px-6 sm:px-10 lg:px-16">
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div
            class="bg-white rounded-3xl p-8 shadow-lg border border-slate-100 hover:-translate-y-2 transition duration-300"
        >
            <div class="text-4xl mb-4">🎵</div>

            <h3 class="text-xl font-bold">Music Events</h3>

            <p class="mt-2 text-slate-500">Discover concerts and live performances.</p>
        </div>

        <div
            class="bg-white rounded-3xl p-8 shadow-lg border border-slate-100 hover:-translate-y-2 transition duration-300"
        >
            <div class="text-4xl mb-4">🏟️</div>

            <h3 class="text-xl font-bold">Sports Events</h3>

            <p class="mt-2 text-slate-500">Find upcoming games and tournaments.</p>
        </div>

        <div
            class="bg-white rounded-3xl p-8 shadow-lg border border-slate-100 hover:-translate-y-2 transition duration-300"
        >
            <div class="text-4xl mb-4">🌍</div>

            <h3 class="text-xl font-bold">Global Experiences</h3>

            <p class="mt-2 text-slate-500">Explore events from cities worldwide.</p>
        </div>
    </div>
</section>

<script src="/static/js/app.js"></script>

{{template "partials/footer.tpl" .}}
