import { Workspace } from '@/modules/auth/domain/user';
import { useAuthStore } from '@/presentation/state-manager/providers/authProvider';
import { Input } from '@mui/material';
import React, { useEffect } from 'react'
import { useShallow } from 'zustand/shallow';

const WorkspaceSelector = () => {
      const { workspaces } = useAuthStore(
        useShallow((state) => ({
          workspaces: state.connectedUser?.workspaces,
        })),
      );


    const onCreateWorkspace = (workspaceName: string) => {

    }

    if(workspaces?.length===0){
        return (
            <div>
                create you first workspace
                <Input/>
            </div>
        )
    }

  return (
    <div>{workspaces?.map(w=>
        <div onClick={onSelectWorkspace(w)}>{w.name}</div>
    )}</div>
  )
}

export default WorkspaceSelector