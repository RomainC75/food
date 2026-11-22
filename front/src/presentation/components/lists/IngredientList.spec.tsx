import { render, screen } from "@testing-library/react";
import { beforeEach} from "node:test";
import IngredientList from "./IngredientList";
import { Providers } from "@/app/provider";
import userEvent from "@testing-library/user-event";
import { assert, describe, it } from "vitest";

describe("ingredientList component", ()=>{
    
    beforeEach(()=>{

    })

    it("should render 2 ingredients", async ()=>{
        Render();
        const res = await screen.findAllByText(/name/)
        assert.equal(res.length, 2)
        // await userEvent.click();
    })
})


const Render = () => {
    render(
        <Providers>
            <IngredientList />
        </Providers>,
    );
};