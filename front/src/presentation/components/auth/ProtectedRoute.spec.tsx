import { cleanup, render, screen } from "@testing-library/react";
import { afterEach, assert, describe, expect, it, vi } from "vitest";
import { IngredientPlain, initPlainIngredients } from "@/modules/recipe/init/ingredient.init";
import Providers from "@/presentation/state-manager/providers";
import { createTestDependencies, TestDependencies } from "@/dependencies";
import ProtectedRoutesComp, { LOGIN_URL } from "./ProtectedRoutes";
import { ConnectedUserInfo } from "@/modules/auth/domain/user";

const pushMock = vi.fn();
const replaceMock = vi.fn(); 

vi.mock("next/navigation", () => ({
  useRouter: () => ({
    push: pushMock,
    replace: replaceMock,
    back: vi.fn(),
    refresh: vi.fn(),
    forward: vi.fn(),
    prefetch: vi.fn(),
  }),
}));

describe("ProtectedRoute component", ()=>{
    afterEach(()=>{
        cleanup();
    })
    
    it("should redirect to '/' if NO authenticated", async()=>{
        Render(initPlainIngredients);
        expect(replaceMock).toHaveBeenCalledWith(LOGIN_URL);
        const text = screen.queryByText("AUTHORIZED")
        expect(text).toBeNull();
        
    })

    it("should display children if authenticated", async ()=>{
        const connectedUser: ConnectedUserInfo = {
            id: "my-id",
            email: "fake@email.com",
            workspaceIds: []
        }
        Render(initPlainIngredients, connectedUser);
        expect(await screen.findByText("AUTHORIZED")).toBeVisible();
    })
})


let deps: TestDependencies;

const Render = (initPlainIngredients: IngredientPlain[], connectedUser?: ConnectedUserInfo) => {
    deps = createTestDependencies({ initPlainIngredients });

    render(
        <Providers deps={deps} connectedUser={connectedUser}>
            <ProtectedRoutesComp>
                <div>AUTHORIZED</div>
            </ProtectedRoutesComp>
        </Providers>,
    );
};