<template>
  <div id="container" class="peer-map-wrapper">

    asdasdasdasdsd
  </div>
</template>

<script>

export default {
  name:"PeerMap",
  data() {
    return {

    }
  },
  mounted() {
    let {am4core, am4charts, am4maps, am4geodata_worldLow} = this.$am4core();
    var chart = am4core.create("container", am4maps.MapChart);
    chart.geodata = am4geodata_worldLow;

    chart.projection = new am4maps.projections.Miller();
    var polygonSeries = chart.series.push(new am4maps.MapPolygonSeries());
    polygonSeries.useGeodata = true;
    polygonSeries.mapPolygons.template.events.on("hit", function(ev) {
      chart.zoomToMapObject(ev.target);
    });
    var label = chart.chartContainer.createChild(am4core.Label);
    label.text = "franceLow";
  }
}
</script>
<style>
  
.peer-map-warpper {
  height: 100% !important;
}

</style>