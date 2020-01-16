<template>
    <apexchart type="radialBar" :options="options" :series="series"></apexchart>
</template>

<script>
    export default {
        data() {
            return {
                series: [10],
                options: {
                    labels: ['CPU']
                }
            };
        },
        // mounted: function() {
        //     wails.events.on("cpu_usage", cpu_usage => {
        //         if (cpu_usage) {
        //             this.series = [ cpu_usage.average ];
        //         }
        //     });
        // }
    };
</script>