import { CircularProgress, Stack, Card, CardContent, Link, Box } from "@mui/joy"
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

      {visibleTranslators.map((translator, index) => {
        return (
          <Card
            key={index}
            orientation="horizontal"
            sx={{
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
              <Link href={translator.details_url} target="_blank" overlay>
                {translator.address}
              </Link>
            </CardContent>
          </Card>
        )
      })}
    </Stack>
  )
}
