import * as am4core from '@amcharts/amcharts4/core'
import * as am4charts from '@amcharts/amcharts4/charts'
import * as am4maps from '@amcharts/amcharts4/maps'
import am4geodata_worldLow from '@amcharts/amcharts4-geodata/worldLow'
import am4themes_kelly from '@amcharts/amcharts4/themes/kelly'
import am4themes_animated from '@amcharts/amcharts4/themes/animated'
// import am4themes_animated from '@amcharts/amcharts4/themes/animated'
// import am4themes_dark from '@amcharts/amcharts4/themes/dark'
import Vue from 'vue'
Vue.prototype.$am4core = () => {
  return {
    am4core,
    am4charts,
    am4maps,
    am4geodata_worldLow,
    am4themes_kelly,
    am4themes_animated,
    // am4themes_animated,
    // am4themes_dark,
  }
}
