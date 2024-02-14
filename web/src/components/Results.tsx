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

export default function Results({ visibleTranslators, loading }: ResultsProps) {
  return (
    <Stack
      spacing={2}
      sx={{
        overflow: "auto",
        px: { xs: 2, md: 4 },
        py: { xs: 4, md: 2 },
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
            rowHeight={130}
            rowRenderer={({ key, index, style }) => {
              const translator = visibleTranslators[index]
              return (
                <Box sx={{ paddingBottom: "10px" }} style={style}>
                  <Card
                    key={key}
                    orientation="horizontal"
                    sx={{
                      height: "120px",
                      bgcolor: "neutral.softBg",
                      display: "flex",
                      flexDirection: { xs: "column", sm: "row" },
                      "&:hover": {
                        boxShadow: "lg",
                        borderColor:
                          "var(--joy-palette-neutral-outlinedDisabledBorder)",
                      },
                    }}
                  >
                    <CardContent>
                      <Stack>
                        <Typography level="title-md">
                          <Link
                            href={translator.details_url}
                            target="_blank"
                            overlay
                          >
                            {translator.name}
                          </Link>
                        </Typography>
                        <Typography>{translator.address}</Typography>
                        <Typography level="body-sm">
                          {translator.location.country}
                        </Typography>
                      </Stack>
                    </CardContent>
                  </Card>
                </Box>
              )
            }}
          />
        )}
      </AutoSizer>
    </Stack>
  )
}
