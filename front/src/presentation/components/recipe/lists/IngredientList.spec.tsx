import { cleanup, render, screen } from "@testing-library/react";
import IngredientList, { NoIngredientFound } from "./IngredientList";
import { afterEach, assert, describe, it } from "vitest";
import { IngredientPlain, initPlainIngredients } from "@/modules/recipe/init/ingredient.init";
import Providers from "@/presentation/state-manager/providers";
import { InMemUuidGen } from "@/modules/shared/infra/in-mem-uuidGen";
import { InMemUser } from "@/modules/auth/infra/in-mem-users";
import { AuthSrv } from "@/modules/auth/app/authSrv";

describe("ingredientList component", ()=>{
    
    afterEach(()=>{
        cleanup();
    })

    it("should render 2 ingredients", async ()=>{
        Render(initPlainIngredients);
        const res = await screen.findAllByText(/name/)
        assert.equal(res.length, 2)
        // await userEvent.click();
    })

    it("should render a message if no ingredient found", async ()=>{
        Render([]);
        const rgx = new RegExp(NoIngredientFound, "")
        const res = await screen.findByText(rgx)
        assert.exists(res)
        // await userEvent.click();
    })
})


const Render = (initPlainIngredients: IngredientPlain[]) => {
    const uuids = new InMemUuidGen()
      const users = new InMemUser(uuids, false)
      const authSrv = new AuthSrv(users)

    render(
        <Providers authSrv= {authSrv} options={{initPlainIngredients}}>
            <IngredientList />
        </Providers>,
    );
};