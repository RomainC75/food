import React from 'react'

const SignupComponent = () => {
        const { ingredients, loading, error } = useRecipeStore(
          useShallow((state) => ({
            ingredients: state.ingredients,
            loading: state.loading,
            error: state.error,
          }))
        );
        const fetchRecipes = useRecipeStore(
          (state) => state.fetchIngredients
        );
  return (
    <div>SignupComponent</div>
  )
}

export default SignupComponent