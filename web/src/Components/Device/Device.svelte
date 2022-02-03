<script>
    export let deviceList;
    $: deviceObj = JSON.parse(deviceList)
    async function generateDash(id) {
        const res = await fetch(location.origin + "/api/device/" + id + "/dash", {
            method: 'POST',
            body: JSON.stringify({"generate": true})
        })
    };
</script>

{#each Object.entries(deviceObj) as [id, dev]}
    <hr>
    <div class="row mb-3">
        <label class="col-sm-2 col-form-label">ID: {id}</label>
        <label class="col-sm-2 col-form-label">Modelo: {dev.type}</label>
        <label class="col-sm-2 col-form-label">Serial: {dev.serial}</label>
        <label class="col-sm-4 col-form-label">Criado: {dev.CreatedAt}</label>
        <div class="col-sm-2">
            <button type="submit" class="btn btn-primary" on:click="{generateDash(id)}">Criar Dash</button>
        </div>
    </div>
{/each}
