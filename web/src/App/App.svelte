<script>
 import Config from '../Components/Config/Config.svelte';
 import { onMount } from 'svelte';
 onMount(async () => {
     console.log(location.href + "api/config");
     fetch(location.href + "api/config")
         .then(response => response.json())
         .then(data => {
		         console.log(data);
             apiData.set(data);
         }).catch(error => {
             console.log(error);
             return [];
         });
 });
</script>

<h1 class="bg-primary text-light text-center">IoT Middleware</h1>

<Config config={"b"}></Config>
