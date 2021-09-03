<template>
  <div>
    <div id="chartdiv">
      asdads
    </div>
  </div>
</template>

<script>

export default {
  name:"PriceChart",
  data() {
    return {
    }
  },
  methods: {
  },
  async mounted() {

    let res = (await this.$axios.get("https://api.binance.com/api/v3/klines?symbol=MINAUSDT&interval=5m&limit=100")).data
    console.log(res)
    let {am4core, am4charts, am4themes_kelly,
    am4themes_animated} = this.$am4core();
      am4core.ready(function() {

      // Themes begin
      am4core.useTheme(am4themes_kelly);
      am4core.useTheme(am4themes_animated);
      // Themes end

      // Create chart instance
      var chart = am4core.create("chartdiv", am4charts.XYChart);
      chart.paddingRight = 20;
      // Add data
      chart.data = res;

      // Create axes
      var categoryAxis = chart.xAxes.push(new am4charts.CategoryAxis());
      categoryAxis.dataFields.category = "0";
      categoryAxis.renderer.minGridDistance = 50;
      categoryAxis.renderer.grid.template.location = 0.5;
      categoryAxis.startLocation = 0.5;
      categoryAxis.endLocation = 0.5;

      // Create value axis
      var valueAxis = chart.yAxes.push(new am4charts.ValueAxis());
      valueAxis.baseValue = 0;

      // Create series
      var series = chart.series.push(new am4charts.LineSeries());
      series.dataFields.valueY = "3";
      series.dataFields.categoryX = "0";
      series.strokeWidth = 2;
      series.tensionX = 0.77;


      // let axisTooltip = categoryAxis.tooltip;

      // axisTooltip.tooltipDateFormat = "yyyy-MM-dd";
      // bullet is added because we add tooltip to a bullet for it to change color
      var bullet = series.bullets.push(new am4charts.Bullet());
      bullet.tooltipText = "{valueY}";
      bullet.tooltipDateFormat = "yyyy-MM-dd";
      bullet.adapter.add("fill", function(fill, target){
          if(target.dataItem.valueY < 0){
              return am4core.color("#FF0000");
          }
          return fill;
      })
      var range = valueAxis.createSeriesRange(series);
      range.value = 0;
      range.endValue = -1000;
      range.contents.stroke = am4core.color("#FF0000");
      range.contents.fill = range.contents.stroke;

      // Add scrollbar
      var scrollbarX = new am4charts.XYChartScrollbar();
      scrollbarX.series.push(series);
      chart.scrollbarX = scrollbarX;

      chart.cursor = new am4charts.XYCursor();
    });
  }
}
</script>
<style>

</style>