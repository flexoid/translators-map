import { Trans } from "@lingui/macro"
import Stack from "@mui/joy/Stack"
import Typography from "@mui/joy/Typography"

export default function HeaderSection() {
  return (
    <Stack sx={{ mb: 2 }}>
      <Stack
        direction="row"
        justifyContent="space-between"
        sx={{ width: "100%" }}
      ></Stack>

      <Typography>
        <Trans id="headerDescription">
          Constantly updated catalog of sworn certified translators of the
          Polish language, compiled from official sources. We offer convenient
          and quick map search, to help you find the perfect specialist for your
          needs.
        </Trans>
      </Typography>
    </Stack>
  )
}
