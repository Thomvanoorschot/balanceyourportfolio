<div id="holdingList">
    <!--{{#each itemList}}-->
    <li class="border-gray-400 flex flex-row mb-2">
        <div class="cursor-pointer bg-gray-200 rounded-md flex flex-1 items-center p-4  transition duration-500 ease-in-out transform hover:-translate-y-1 hover:shadow-lg">
            <div class="flex flex-col rounded-md w-10 h-10 bg-gray-300 justify-center items-center mr-4">
                <!--{{#if (equals this.Type 'Currency')}}-->
                <!--$-->
                <!--{{else}}-->
                <img src="/company-logos/AAPL.png" onerror="this.onerror=null; this.src='/company-logos/UNKNOWN.png'" >
                <!--{{/if}}-->
            </div>
            <div class="flex-1 pl-1 mr-16">
                <div class="font-medium">Ticker</div>
                <div class="text-gray-600 text-sm">Name</div>
            </div>
            <div class="text-gray-600 text-xs">PerentageOfTotal</div>
        </div>
    </li>
    <!--{{/each}}-->
</div>
