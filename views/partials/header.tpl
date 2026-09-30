<!doctype html>

<html lang="en">
    <head>
        <meta charset="UTF-8" />

        <meta name="viewport" content="width=device-width, initial-scale=1.0" />

        <title>Event Explorer</title>

        <!-- Google Font -->
        <link rel="preconnect" href="https://fonts.googleapis.com" />

        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />

        <link
            href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap"
            rel="stylesheet"
        />

        <!-- Tailwind CDN -->

        <script src="https://cdn.tailwindcss.com"></script>

        <link rel="stylesheet" href="/static/css/style.css" />

        <script>
            tailwind.config = {
                theme: {
                    extend: {
                        fontFamily: {
                            sans: ["Inter", "Arial", "sans-serif"],
                        },

                        colors: {
                            brand: {
                                500: "#2563eb",

                                600: "#1d4ed8",
                            },
                        },
                    },
                },
            };
        </script>

        <style>
            html {
                scroll-behavior: smooth;
            }

            body {
                font-family: Inter, sans-serif;
            }

            .fade-in {
                animation: fadeIn 0.5s ease-in-out;
            }

            @keyframes fadeIn {
                from {
                    opacity: 0;

                    transform: translateY(10px);
                }

                to {
                    opacity: 1;

                    transform: translateY(0);
                }
            }
        </style>
    </head>

    <body class="bg-slate-50 text-slate-900">
        <header class="sticky top-0 z-50 bg-white/80 backdrop-blur-lg border-b border-slate-200">
            <div class="max-w-7xl mx-auto px-6 py-5 flex items-center justify-between">
                <!-- Logo -->

                <a href="/" class="flex items-center gap-3 group">
                    <div
                        class="w-11 h-11 rounded-2xl bg-gradient-to-br from-blue-600 to-indigo-600 flex items-center justify-center text-white text-xl shadow-lg group-hover:scale-105 transition"
                    >
                        🎟
                    </div>

                    <div>
                        <h1 class="text-xl font-extrabold tracking-tight text-slate-900">Event Explorer</h1>

                        <p class="text-xs text-slate-500">Discover experiences</p>
                    </div>
                </a>

                <!-- Navigation -->

                <nav class="hidden md:flex items-center gap-8 text-sm font-medium">
                    <a href="/" class="text-slate-600 hover:text-blue-600 transition"> Home </a>

                    <a href="/" class="text-slate-600 hover:text-blue-600 transition"> Explore </a>

                    <a href="/" class="text-slate-600 hover:text-blue-600 transition"> About </a>
                </nav>
            </div>
        </header>

        <main class="max-w-7xl mx-auto px-6 sm:px-10 lg:px-16 xl:px-20 py-12 fade-in"></main>
    </body>
</html>
