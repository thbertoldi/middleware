<script>
 import { onMount } from 'svelte';
 import Device from '../Components/Device/Device.svelte'
 import Config from '../Components/Config/Config.svelte'
 $: appConfig = {"http": {"port": 0}, "mqtt": {"host": "", "port": 0, "topic": [], "id": ""}};
 $: allowJoin = true;
 $: deviceList = {};
 onMount(async () => {
     fetch(location.origin + "/api/config")
         .then(response => response.json())
         .then(data => {
             appConfig = data
         }).catch(error => {
             return [];
         });
     fetch(location.origin + "/api/device")
         .then(response => response.json())
         .then(data => {
             deviceList = data;
         }).catch(error => {
             return [];
         });
 });
</script>

<html lang="en">
    <body class="bg-light">
        <h1 class="bg-secondary text-light text-center">IoT Middleware</h1>
        <div class="container">
            <Config {appConfig} {allowJoin}></Config>
            <Device {deviceList}></Device>
        </div>
    </body>
</html>
