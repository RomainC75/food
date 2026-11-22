"use client"
import { useEffect } from "react";
import { useRecipeStore } from "./provider";
import { CreateIngredientFormComponent } from "@/presentation/components/inputs/createIngredient/CreateIngredient";

export default function Home() {
    const ingredients = useRecipeStore((state) => state.ingredients);
  const loading = useRecipeStore((state) => state.loading);
  const error = useRecipeStore((state) => state.error);
  const fetchRecipes = useRecipeStore(
    (state) => state.fetchIngredients
  );

  useEffect(() => {
    console.log('-- useEffect')
    fetchRecipes();

  }, [fetchRecipes]);

  if (loading) return <p>Loading recipes...</p>;

  if (error) return <p>Error: {error}</p>;

  return (
    <div className="flex flex-col flex-1 items-center justify-center bg-zinc-50 font-sans dark:bg-black">
      <h1>Ingredients</h1>
      <CreateIngredientFormComponent/>
      <main className="flex flex-1 w-full max-w-3xl flex-col items-center justify-between py-32 px-16 bg-white dark:bg-black sm:items-start">
        {ingredients.map((ing, index)=>
          <div key={index}>
            <div> id : {ing.Id}</div>
            <div> name : {ing.Name}</div>
            <div> unit : {ing.BaseUnit}</div>
          </div>
        )}
      </main>
    </div>
  );
}
