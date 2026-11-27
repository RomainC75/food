import { Box, CircularProgress } from "@mui/material";

const FullPageSpinner = () => (
  <Box
    role="status"
    aria-label="Loading"
    sx={{
      display: "flex",
      alignItems: "center",
      justifyContent: "center",
      minHeight: "100vh",
      width: "100%",
    }}
  >
    <CircularProgress />
  </Box>
);

export default FullPageSpinner;
