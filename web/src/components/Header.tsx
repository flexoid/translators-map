import { Box, IconButton, Typography } from "@mui/joy"
import PublicIcon from "@mui/icons-material/Public"
import { Trans } from "@lingui/macro"

export default function Header() {
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "row",
        justifyContent: "space-between",
        alignItems: "center",
        width: "100%",
        top: 0,
        px: 1.5,
        py: 1,
        zIndex: 10000,
        backgroundColor: "background.body",
        borderBottom: "1px solid",
        borderColor: "divider",
        position: "sticky",
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          gap: 1.5,
        }}
      >
        <IconButton size="sm" variant="soft">
          <PublicIcon />
        </IconButton>
        <Typography component="h1" fontWeight="xl">
          <Trans id="title">Sworn Translator Map</Trans> -{" "}
          <Trans id="subtitle">
            Your directory of Polish Sworn Translators
          </Trans>
        </Typography>
      </Box>
    </Box>
  )
}
