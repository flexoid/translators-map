import { useRef, useState, useEffect } from "react"
import { MarkerClusterer } from "@googlemaps/markerclusterer"
import { Translator } from "../lib/api"

interface MapProps extends google.maps.MapOptions {
  translators: Translator[]
  style: { [key: string]: string }
  onClick?: (e: google.maps.MapMouseEvent) => void
  onIdle?: (map: google.maps.Map) => void
}

function Map({ center, zoom, style, translators, ...options }: MapProps) {
  const ref = useRef<HTMLDivElement>(null)
  const [map, setMap] = useState<google.maps.Map>()
  const [clusterer, setClusterer] = useState<any>(null)
  const [infoWindow, setInfoWindow] = useState<google.maps.InfoWindow | null>(
    null
  )

  useEffect(() => {
    if (ref.current && !map) {
      const map = new google.maps.Map(ref.current, { center, zoom })
      setClusterer(new MarkerClusterer({ map }))
      setMap(map)
      setInfoWindow(new google.maps.InfoWindow())
    }
  }, [ref, map])

  useEffect(() => {
    if (!map || !translators || !infoWindow) {
      return
    }

    const markers = translators.map((translator) => {
      const marker = new google.maps.Marker({
        position: new google.maps.LatLng(
          translator.location.lat,
          translator.location.lng
        ),
      })

      marker.addListener("click", () => {
        const infoContent = `
          <div class="info-window">
            <b>${translator.name}</b><br>${translator.address}<br>${translator.contacts}<br>
            <a href="${translator.details_url}" target="_blank">Open details in BIP (new tab)</a>
          </div>
        `

        infoWindow.setContent(infoContent)
        infoWindow.open(map, marker)
      })

      return marker
    })

    if (clusterer) {
      clusterer.clearMarkers()
      clusterer.addMarkers(markers)
    }
  }, [map, translators])

  return <div ref={ref} style={style} />
}

export default Map
