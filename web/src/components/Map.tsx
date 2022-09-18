import { useRef, useState, useEffect } from "react"
import { MarkerClusterer } from "@googlemaps/markerclusterer"
import { renderToString } from "react-dom/server"
import { ExternalLinkIcon } from "@chakra-ui/icons"
import { Translator } from "../lib/api"

interface MapProps extends google.maps.MapOptions {
  translators: Translator[]
  style: { [key: string]: string }
  onClick?: (e: google.maps.MapMouseEvent) => void
  onVisibleTranslatorsChange?: (translators: Translator[]) => void
}

interface MarkerItem {
  translator: Translator
  marker: google.maps.Marker
}

function MapComponent({
  center,
  zoom,
  style,
  translators,
  onVisibleTranslatorsChange,
  ...options
}: MapProps) {
  const ref = useRef<HTMLDivElement>(null)
  const [map, setMap] = useState<google.maps.Map>()
  const [clusterer, setClusterer] = useState<MarkerClusterer | null>(null)
  const [infoWindow, setInfoWindow] = useState<google.maps.InfoWindow | null>(
    null
  )
  const markersRef = useRef<Map<google.maps.Marker, Translator>>(new Map())

  const externalLinkIconStr = renderToString(<ExternalLinkIcon mx="2px" />)

  function updateVisibleTranslators(map: google.maps.Map) {
    const markers = markersRef.current

    let bounds = map.getBounds()
    if (!bounds) return

    const visibleTranslators = Array.from(markers)
      .filter(([marker, translator]) => {
        let pos = marker.getPosition()
        if (pos && bounds!.contains(pos)) {
          return true
        }
      })
      .map(([marker, translator]) => translator)

    const center = map.getCenter()
    if (!center) return

    visibleTranslators.sort((a, b) => {
      const aToCenter = google.maps.geometry.spherical.computeDistanceBetween(
        center,
        new google.maps.LatLng(a.location.lat, a.location.lng)
      )

      const bToCenter = google.maps.geometry.spherical.computeDistanceBetween(
        center,
        new google.maps.LatLng(b.location.lat, b.location.lng)
      )

      return bToCenter - aToCenter
    })

    onVisibleTranslatorsChange?.(visibleTranslators)
  }

  useEffect(() => {
    if (ref.current && !map) {
      const map = new google.maps.Map(ref.current, { center, zoom })
      setClusterer(new MarkerClusterer({ map }))
      setMap(map)
      setInfoWindow(new google.maps.InfoWindow())

      google.maps.event.addListener(map, "idle", function () {
        updateVisibleTranslators(map)
      })
    }
  }, [ref, map])

  useEffect(() => {
    if (!map || !translators || !infoWindow) {
      return
    }

    const markerItems = new Map(
      translators.map((translator) => {
        const marker = new google.maps.Marker({
          position: new google.maps.LatLng(
            translator.location.lat,
            translator.location.lng
          ),
        })

        marker.addListener("click", () => {
          const infoContent = `
          <div class="info-window">
            <a href="${translator.details_url}" target="_blank">See details ${externalLinkIconStr}</a>
          </div>
        `

          infoWindow.setContent(infoContent)
          infoWindow.open(map, marker)
        })

        return [marker, translator]
      })
    )

    if (clusterer) {
      clusterer.clearMarkers()
      clusterer.addMarkers(Array.from(markerItems.keys()))
      markersRef.current = markerItems
      updateVisibleTranslators(map)
    }
  }, [map, translators])

  return <div ref={ref} style={style} />
}

export default MapComponent
