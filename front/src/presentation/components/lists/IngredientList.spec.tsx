import { render, screen } from "@testing-library/react";
import { beforeEach} from "node:test";
import IngredientList, { NoIngredientFound } from "./IngredientList";
import { Providers } from "@/app/provider";
import userEvent from "@testing-library/user-event";
import { assert, describe, it } from "vitest";
import { initIngredients } from "@/modules/recipe/infra/in-mem-ingredients";
import { Ingredient } from "@/modules/recipe/domain/Ingredient";

describe("ingredientList component", ()=>{
    
    beforeEach(()=>{

    })

    it("should render 2 ingredients", async ()=>{
        Render(initIngredients);
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


const Render = (initIngredients: Ingredient[]) => {
    render(
        <Providers initIngredients={initIngredients}>
            <IngredientList />
        </Providers>,
    );
};