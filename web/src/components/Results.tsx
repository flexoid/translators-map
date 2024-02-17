import { AutoSizer, List } from "react-virtualized"
import {
  CircularProgress,
  Stack,
  Card,
  CardContent,
  Link,
  Box,
  Typography,
} from "@mui/joy"
import { Translator } from "../lib/api"

type ResultsProps = {
  visibleTranslators: Translator[]
  loading: boolean
}

function TranslatorItem(translator: Translator) {
  return (
    <Card
      orientation="horizontal"
      sx={{
        height: "100%",
        bgcolor: "neutral.softBg",
        display: "flex",
        flexDirection: { xs: "column", sm: "row" },
        "&:hover": {
          boxShadow: "lg",
          borderColor: "var(--joy-palette-neutral-outlinedDisabledBorder)",
        },
      }}
    >
      <CardContent>
        <Stack>
          <Typography level="title-sm">
            <Link href={translator.details_url} target="_blank" overlay>
              {translator.name}
            </Link>
          </Typography>
          <Typography>{translator.address}</Typography>
          <Typography level="body-sm">{translator.location.country}</Typography>
        </Stack>
      </CardContent>
    </Card>
  )
}

export default function Results({ visibleTranslators, loading }: ResultsProps) {
  return (
    <Stack
      spacing={2}
      sx={{
        px: { xs: 2, md: 0 },
        py: 2,
        height: { xs: "40vh", md: "auto" },
      }}
    >
      <Box sx={{ display: "flex", justifyContent: "center" }}>
        {loading && <CircularProgress size="lg" />}
      </Box>

      <AutoSizer>
        {({ height, width }) => (
          <List
            width={width}
            height={height}
            rowCount={visibleTranslators.length}
            rowHeight={105}
            rowRenderer={({ key, index, style }) => {
              const translator = visibleTranslators[index]
              return (
                <Box
                  sx={{ height: "100px", paddingBottom: "5px" }}
                  style={style}
                  key={key}
                >
                  <TranslatorItem {...translator} />
                </Box>
              )
            }}
          />
        )}
      </AutoSizer>
    </Stack>
  )
}
