<script>
    import "../app.css";
</script>


<div id="container" class="flex flex-col min-h-screen">
    <header class="flex justify-between p-3 items-center">
        <div>LOGO</div>
        <div class="flex items-center">
            <a href="search" class="">Search</a>
            <a href="portfolios" class="">Portfolios</a>
        </div>
        <div class="flex items-center">
            <div class="p-2 rounded hover:bg-sky-200">Login</div>
            <div class="ml-3 border-solid border-2 rounded border-sky-500 p-2 hover:bg-sky-200" >Signup</div>
        </div>
    </header>

    <slot/>
    
    <footer class="bg-white text-center">
        <div class="p-4 text-center text-neutral-800">
            © 2023 Copyright:
            <a class="text-neutral-800" href="https://tailwind-elements.com/">EtfInsight</a>
        </div>
    </footer>
</div>
