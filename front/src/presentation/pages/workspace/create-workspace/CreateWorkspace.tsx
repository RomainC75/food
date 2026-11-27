"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useShallow } from "zustand/shallow";
import {
  Dialog,
  DialogContent,
  Box,
  Typography,
  TextField,
  Button,
  Stack,
  Avatar,
  Alert,
  CircularProgress,
} from "@mui/material";
import WorkspacesOutlinedIcon from "@mui/icons-material/WorkspacesOutlined";
import ArrowForwardRoundedIcon from "@mui/icons-material/ArrowForwardRounded";
import { useAuthStore } from "@/presentation/state-manager/providers/authProvider";
import {
  WORKSPACE_NAME_MAX as NAME_MAX,
  WORKSPACE_NAME_MIN as NAME_MIN,
} from "@/modules/auth/app/workspaceSrv";

export default function CreateWorkspace() {
  const router = useRouter();
  const [workspaceName, setWorkspaceName] = useState("");
  const [touched, setTouched] = useState(false);
  const { createWorkspace, loading, error } = useAuthStore(
    useShallow((state) => ({
      createWorkspace: state.createWorkspace,
      loading: state.createWorkspacePending,
      error: state.createWorkspaceError,
    })),
  );

  const name = workspaceName.trim();
  const isValid = name.length >= NAME_MIN && name.length <= NAME_MAX;
  const showLengthError = touched && !isValid;

  const handleCreate = async () => {
    setTouched(true);
    if (!isValid || loading) return;

    const ok = await createWorkspace(name);
    if (ok) router.push("/");
  };

  return (
    <Dialog
      open={true}
      fullWidth
      maxWidth="xs"
      slotProps={{
        backdrop: {
          sx: {
            backgroundColor: "rgba(15, 23, 42, 0.65)",
            backdropFilter: "blur(5px)",
          },
        },
        paper: {
          sx: {
            borderRadius: 4,
            p: 2,
            boxShadow: "0 24px 80px rgba(0,0,0,0.2)",
          },
        },
      }}
    >
      <DialogContent>
        <Stack
          component="form"
          spacing={3}
          sx={{ alignItems: "center" }}
          noValidate
          onSubmit={(e) => {
            e.preventDefault();
            handleCreate();
          }}
        >
          <Avatar
            sx={{
              width: 64,
              height: 64,
              bgcolor: "#EEF2FF",
              color: "#4F46E5",
            }}
          >
            <WorkspacesOutlinedIcon fontSize="large" />
          </Avatar>

          <Box sx={{ textAlign: "center" }}>
            <Typography variant="h5" sx={{ fontWeight: 700 }}>
              Create your workspace
            </Typography>

            <Typography variant="body2" sx={{ color: "text.secondary", mt: 1 }}>
              Give your workspace a name to get started. You can always change
              it later.
            </Typography>
          </Box>

          <Box sx={{ width: "100%" }}>
            <Typography
              component="label"
              htmlFor="workspace-name"
              variant="body2"
              sx={{ display: "block", mb: 1, fontWeight: 600 }}
            >
              Workspace name
            </Typography>

            <TextField
              id="workspace-name"
              fullWidth
              autoFocus
              placeholder="e.g. Acme Studio"
              value={workspaceName}
              disabled={loading}
              onChange={(e) => {
                setWorkspaceName(e.target.value);
              }}
              onBlur={() => setTouched(true)}
              error={showLengthError}
              helperText={
                showLengthError
                  ? `Between ${NAME_MIN} and ${NAME_MAX} characters`
                  : `${name.length}/${NAME_MAX}`
              }
              slotProps={{ htmlInput: { maxLength: NAME_MAX } }}
            />
          </Box>

          {error && (
            <Alert severity="error" sx={{ width: "100%" }}>
              {error}
            </Alert>
          )}

          <Button
            type="submit"
            fullWidth
            variant="contained"
            size="large"
            disabled={!isValid || loading}
            endIcon={
              loading ? (
                <CircularProgress size={18} color="inherit" />
              ) : (
                <ArrowForwardRoundedIcon />
              )
            }
            sx={{
              py: 1.5,
              borderRadius: 2,
              textTransform: "none",
              fontWeight: 600,
              bgcolor: "#4F46E5",
              "&:hover": { bgcolor: "#4338CA" },
            }}
          >
            {loading ? "Creating…" : "Create workspace"}
          </Button>

          <Typography variant="caption" sx={{ color: "text.secondary" }}>
            Your workspace is where your team and projects live.
          </Typography>
        </Stack>
      </DialogContent>
    </Dialog>
  );
}
