<template>
  <div id="chartdiv">
  </div>
</template>

<script>

export default {
  name:"PeerMap",
  data() {
    return {

    }
  },
  methods: {

    // this function creates and returns a new marker element

  },
  async mounted() {

    const res = (await this.$axios.get("/getPeers")).data.Data

    let {am4core, am4charts, am4maps, am4geodata_worldLow} = this.$am4core();
    am4core.ready(() => {

      var chart = am4core.create("chartdiv", am4maps.MapChart);
      chart.geodata = am4geodata_worldLow;
      chart.projection = new am4maps.projections.Miller();
      var polygonSeries = chart.series.push(new am4maps.MapPolygonSeries());
      polygonSeries.exclude = ["AQ"];
      polygonSeries.useGeodata = true;

      var polygonTemplate = polygonSeries.mapPolygons.template;
      polygonTemplate.tooltipText = "{name}";
      polygonTemplate.fill = chart.colors.getIndex(0).lighten(0.5);

      var hs = polygonTemplate.states.create("hover");
      hs.properties.fill = chart.colors.getIndex(0);

      var imageSeries = chart.series.push(new am4maps.MapImageSeries());
      imageSeries.mapImages.template.propertyFields.longitude = "longitude";
      imageSeries.mapImages.template.propertyFields.latitude = "latitude";
      imageSeries.data = res
      chart.events.on( "ready", updateCustomMarkers );
      chart.events.on( "mappositionchanged", updateCustomMarkers );

      function updateCustomMarkers( event ) {
        imageSeries.mapImages.each(function(image) {
          if (!image.dummyData || !image.dummyData.externalElement) {
            image.dummyData = {
              externalElement: createCustomMarker(image)
            };
          }

          var xy = chart.geoPointToSVG( { longitude: image.longitude, latitude: image.latitude } );
          image.dummyData.externalElement.style.top = xy.y + 'px';
          image.dummyData.externalElement.style.left = xy.x + 'px';
        });

      }

      function createCustomMarker( image ) {
        
        var chart = image.dataItem.component.chart;

      var holder = document.createElement( 'div' );
        holder.className = 'map-marker';
        holder.title = image.dataItem.dataContext.title;
        holder.style.position = 'absolute';

        if ( undefined != image.url ) {
          holder.onclick = function() {
            window.location.href = image.url;
          };
          holder.className += ' map-clickable';
        }

        var dot = document.createElement( 'div' );
        dot.className = 'dot';
        holder.appendChild( dot );

        var pulse = document.createElement( 'div' );
        pulse.className = 'pulse';
        holder.appendChild( pulse );

        chart.svgContainer.htmlElement.appendChild( holder );

        return holder;
      }

    })
  }
}
</script>
<style>
  
#chartdiv {
  width: auto;
  height: 500px !important;
  margin: 15px;
  overflow: hidden;
  position: relative;
  border-radius: 15px;
  border: 1px solid grey
}


.map-marker {
  margin-left: -8px;
  margin-top: -8px;
  box-sizing: border-box;
}
.map-marker.map-clickable {
  cursor: pointer;
}
.pulse {
  width: 10px;
  height: 10px;
  border: 5px solid #f7f14c;
  -webkit-border-radius: 30px;
  -moz-border-radius: 30px;
  border-radius: 30px;
  background-color: #716f42;
  z-index: 10;
  position: absolute;
  box-sizing: border-box;
}
.map-marker .dot {
  border: 10px solid #fff601;
  background: transparent;
  -webkit-border-radius: 60px;
  -moz-border-radius: 60px;
  border-radius: 60px;
  height: 50px;
  width: 50px;
  -webkit-animation: pulse 3s ease-out;
  -moz-animation: pulse 3s ease-out;
  animation: pulse 3s ease-out;
  -webkit-animation-iteration-count: infinite;
  -moz-animation-iteration-count: infinite;
  animation-iteration-count: infinite;
  position: absolute;
  top: -20px;
  left: -20px;
  z-index: 1;
  opacity: 0;
  box-sizing: border-box;
}

@-moz-keyframes pulse {
  0% {
     -moz-transform: scale(0);
     opacity: 0.0;
  }
  25% {
     -moz-transform: scale(0);
     opacity: 0.1;
  }
  50% {
     -moz-transform: scale(0.1);
     opacity: 0.3;
  }
  75% {
     -moz-transform: scale(0.5);
     opacity: 0.5;
  }
  100% {
     -moz-transform: scale(1);
     opacity: 0.0;
  }
}

@-webkit-keyframes pulse {
  0% {
    -webkit-transform: scale(0);
    opacity: 0.0;
  }
  25% {
    -webkit-transform: scale(0);
    opacity: 0.1;
  }
  50% {
    -webkit-transform: scale(0.1);
    opacity: 0.3;
  }
  75% {
    -webkit-transform: scale(0.5);
    opacity: 0.5;
  }
  100% {
    -webkit-transform: scale(1);
    opacity: 0.0;
  }
}
</style>