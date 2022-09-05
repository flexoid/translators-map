function initMap() {
  const map = new google.maps.Map(document.getElementById("map"), {
    zoom: 7,
    center: { lat: 52.237049, lng: 21.017532 },
  });

  let infowindow = new google.maps.InfoWindow();

  // TODO: Make configurable.
  const lang = "białoruski";

  fetch(`/api/translators?lang=${lang}`)
    .then((response) => {
      return response.json();
    })
    .then((data) => {
      const markers = data.map((translator, i) => {
        const marker = new google.maps.Marker({
          position: new google.maps.LatLng(translator.location.lat, translator.location.lng),
        });

        marker.addListener("click", () => {
          let infoContent = `<b>${translator.name}</b><br>${translator.address}<br>${translator.contacts}`
          infowindow.setContent(infoContent);
          infowindow.open(map, marker);
        });

        return marker;
      });

      new markerClusterer.MarkerClusterer({ map, markers });
    });
}

window.initMap = initMap;

fetch("/api/config")
  .then((response) => {
    return response.json();
  })
  .then((config) => {
    const script = document.createElement('script');
    script.src = `https://maps.googleapis.com/maps/api/js?key=${config.maps_js_api_key}&callback=initMap`;
    script.async = true;

    document.head.appendChild(script);
  });
