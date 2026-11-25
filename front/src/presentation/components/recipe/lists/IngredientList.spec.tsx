import { cleanup, render, screen } from "@testing-library/react";
import IngredientList, { NoIngredientFound } from "./IngredientList";
import { afterEach, assert, describe, it } from "vitest";
import { IngredientPlain, initPlainIngredients } from "@/modules/recipe/init/ingredient.init";
import Providers from "@/presentation/state-manager/providers";
import { createTestDependencies, TestDependencies } from "@/dependencies";

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


let deps: TestDependencies;

const Render = (initPlainIngredients: IngredientPlain[]) => {
    deps = createTestDependencies({ initPlainIngredients });

    render(
        <Providers deps={deps}>
            <IngredientList />
        </Providers>,
    );
};