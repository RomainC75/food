import { cleanup, render, screen } from "@testing-library/react";
import IngredientList, { NoIngredientFound } from "./IngredientList";
import { afterEach, assert, describe, it } from "vitest";
import { IngredientPlain, initPlainIngredients } from "@/modules/recipe/init/ingredient.init";
import Providers from "@/presentation/state-manager/providers";

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
    render(
        <Providers options={{initPlainIngredients}}>
            <IngredientList />
        </Providers>,
    );
};