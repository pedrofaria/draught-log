<script>
    import AttrItem from "./AttrItem.svelte";

    export let payload;
    export let prefix;

    function transformAttrDotToObject(data) {
        let transformed = {}

        Object.keys(data).forEach((attr) => {
            const val = data[attr]
            const keys = attr.split('.')

            if (keys.length === 1) {
                if (typeof val === "object") {
                    transformed[attr] = transformAttrDotToObject(val)
                } else {
                    transformed[attr] = data[attr];
                }
                return;
            }

            const k = keys.shift();
            const newKeys = keys.join('.')
            let nData = {};

            if (transformed.hasOwnProperty(k)) {
                nData[k] = transformed[k];
            } else {
                nData[k] = {}
            }

            nData[k][newKeys] = val;

            nData[k] = transformAttrDotToObject(nData[k])

            transformed = {...transformed, ...nData}
        })

        return transformed
    }

    $: preparedPayload = transformAttrDotToObject(payload);
</script>

<table>
    <tbody>
    {#each Object.keys(preparedPayload) as attr}
        <AttrItem prefix={prefix} key={attr} value={preparedPayload[attr]} />
    {/each}
    </tbody>
</table>
