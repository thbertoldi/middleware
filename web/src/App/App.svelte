<script>
 import { onMount } from 'svelte';
 import Device from '../Components/Device/Device.svelte'
 import Config from '../Components/Config/Config.svelte'
 $: appConfig = {"http": {"port": 0}, "mqtt": {"host": "", "port": 0, "topic": [], "id": ""}};
 $: allowJoin = true;
 $: deviceList = {};
 $: cfg = true;
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

<svelte:head>
    <title>IoT Middleware</title>
</svelte:head>

<html lang="en">
    <body class="bg-light">
    <div class="container-fluid">
        <div class="row flex-nowrap">
            <div class="col-auto px-0">
                <div id="sidebar" class="collapse collapse-horizontal show border-end">
                    <div id="sidebar-nav" class="list-group border-0 rounded-0 text-sm-start min-vh-100">
                        <a href="#" on:click="{() => (cfg = true)}" class="list-group-item border-end-0 d-inline-block text-truncate" data-bs-parent="#sidebar"><span>Configurações</span> </a>
                        <a href="#" on:click="{() => (cfg = false)}" class="list-group-item border-end-0 d-inline-block text-truncate" data-bs-parent="#sidebar"><span>Dispositivos</span> </a>
                    </div>
                </div>
            </div>
            <main class="col ps-md-2 pt-2">
                <a href="#" data-bs-target="#sidebar" data-bs-toggle="collapse" class="border rounded-3 p-1 text-decoration-none"><i class="bi bi-list bi-lg py-2 p-1"></i> Menu</a>
                <div class="row">
                    <div class="col-12">
                        <h1 class="bg-secondary text-light text-center">IoT Middleware</h1>
                        <div class="container">
                            {#if cfg}
                                <Config {appConfig} {allowJoin}></Config>
                            {:else}
                                <Device {deviceList}></Device>
                            {/if}
                        </div>
                    </div>
                </div>
            </main>
        </div>
    </div>
    </body>
</html>
