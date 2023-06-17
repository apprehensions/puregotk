// Package gsk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gsk

// The blend modes available for render nodes.
//
// The implementation of each blend mode is deferred to the
// rendering pipeline.
//
// See &lt;https://www.w3.org/TR/compositing-1/#blending&gt; for more information
// on blending and blend modes.
type BlendMode int

const (

	// The default blend mode, which specifies no blending
	BlendModeDefaultValue BlendMode = 0
	// The source color is multiplied by the destination
	//   and replaces the destination
	BlendModeMultiplyValue BlendMode = 1
	// Multiplies the complements of the destination and source
	//   color values, then complements the result.
	BlendModeScreenValue BlendMode = 2
	// Multiplies or screens the colors, depending on the
	//   destination color value. This is the inverse of hard-list
	BlendModeOverlayValue BlendMode = 3
	// Selects the darker of the destination and source colors
	BlendModeDarkenValue BlendMode = 4
	// Selects the lighter of the destination and source colors
	BlendModeLightenValue BlendMode = 5
	// Brightens the destination color to reflect the source color
	BlendModeColorDodgeValue BlendMode = 6
	// Darkens the destination color to reflect the source color
	BlendModeColorBurnValue BlendMode = 7
	// Multiplies or screens the colors, depending on the source color value
	BlendModeHardLightValue BlendMode = 8
	// Darkens or lightens the colors, depending on the source color value
	BlendModeSoftLightValue BlendMode = 9
	// Subtracts the darker of the two constituent colors from the lighter color
	BlendModeDifferenceValue BlendMode = 10
	// Produces an effect similar to that of the difference mode but lower in contrast
	BlendModeExclusionValue BlendMode = 11
	// Creates a color with the hue and saturation of the source color and the luminosity of the destination color
	BlendModeColorValue BlendMode = 12
	// Creates a color with the hue of the source color and the saturation and luminosity of the destination color
	BlendModeHueValue BlendMode = 13
	// Creates a color with the saturation of the source color and the hue and luminosity of the destination color
	BlendModeSaturationValue BlendMode = 14
	// Creates a color with the luminosity of the source color and the hue and saturation of the destination color
	BlendModeLuminosityValue BlendMode = 15
)

// The corner indices used by `GskRoundedRect`.
type Corner int

const (

	// The top left corner
	CornerTopLeftValue Corner = 0
	// The top right corner
	CornerTopRightValue Corner = 1
	// The bottom right corner
	CornerBottomRightValue Corner = 2
	// The bottom left corner
	CornerBottomLeftValue Corner = 3
)

// This defines the types of the uniforms that `GskGLShaders`
// declare.
//
// It defines both what the type is called in the GLSL shader
// code, and what the corresponding C type is on the Gtk side.
type GLUniformType int

const (

	// No type, used for uninitialized or unspecified values.
	GlUniformTypeNoneValue GLUniformType = 0
	// A float uniform
	GlUniformTypeFloatValue GLUniformType = 1
	// A GLSL int / gint32 uniform
	GlUniformTypeIntValue GLUniformType = 2
	// A GLSL uint / guint32 uniform
	GlUniformTypeUintValue GLUniformType = 3
	// A GLSL bool / gboolean uniform
	GlUniformTypeBoolValue GLUniformType = 4
	// A GLSL vec2 / graphene_vec2_t uniform
	GlUniformTypeVec2Value GLUniformType = 5
	// A GLSL vec3 / graphene_vec3_t uniform
	GlUniformTypeVec3Value GLUniformType = 6
	// A GLSL vec4 / graphene_vec4_t uniform
	GlUniformTypeVec4Value GLUniformType = 7
)

// The type of a node determines what the node is rendering.
type RenderNodeType int

const (

	// Error type. No node will ever have this type.
	NotARenderNodeValue RenderNodeType = 0
	// A node containing a stack of children
	ContainerNodeValue RenderNodeType = 1
	// A node drawing a `cairo_surface_t`
	CairoNodeValue RenderNodeType = 2
	// A node drawing a single color rectangle
	ColorNodeValue RenderNodeType = 3
	// A node drawing a linear gradient
	LinearGradientNodeValue RenderNodeType = 4
	// A node drawing a repeating linear gradient
	RepeatingLinearGradientNodeValue RenderNodeType = 5
	// A node drawing a radial gradient
	RadialGradientNodeValue RenderNodeType = 6
	// A node drawing a repeating radial gradient
	RepeatingRadialGradientNodeValue RenderNodeType = 7
	// A node drawing a conic gradient
	ConicGradientNodeValue RenderNodeType = 8
	// A node stroking a border around an area
	BorderNodeValue RenderNodeType = 9
	// A node drawing a `GdkTexture`
	TextureNodeValue RenderNodeType = 10
	// A node drawing an inset shadow
	InsetShadowNodeValue RenderNodeType = 11
	// A node drawing an outset shadow
	OutsetShadowNodeValue RenderNodeType = 12
	// A node that renders its child after applying a matrix transform
	TransformNodeValue RenderNodeType = 13
	// A node that changes the opacity of its child
	OpacityNodeValue RenderNodeType = 14
	// A node that applies a color matrix to every pixel
	ColorMatrixNodeValue RenderNodeType = 15
	// A node that repeats the child's contents
	RepeatNodeValue RenderNodeType = 16
	// A node that clips its child to a rectangular area
	ClipNodeValue RenderNodeType = 17
	// A node that clips its child to a rounded rectangle
	RoundedClipNodeValue RenderNodeType = 18
	// A node that draws a shadow below its child
	ShadowNodeValue RenderNodeType = 19
	// A node that blends two children together
	BlendNodeValue RenderNodeType = 20
	// A node that cross-fades between two children
	CrossFadeNodeValue RenderNodeType = 21
	// A node containing a glyph string
	TextNodeValue RenderNodeType = 22
	// A node that applies a blur
	BlurNodeValue RenderNodeType = 23
	// Debug information that does not affect the rendering
	DebugNodeValue RenderNodeType = 24
	// A node that uses OpenGL fragment shaders to render
	GlShaderNodeValue RenderNodeType = 25
)

// The filters used when scaling texture data.
//
// The actual implementation of each filter is deferred to the
// rendering pipeline.
type ScalingFilter int

const (

	// linear interpolation filter
	ScalingFilterLinearValue ScalingFilter = 0
	// nearest neighbor interpolation filter
	ScalingFilterNearestValue ScalingFilter = 1
	// linear interpolation along each axis,
	//   plus mipmap generation, with linear interpolation along the mipmap
	//   levels
	ScalingFilterTrilinearValue ScalingFilter = 2
)

// Errors that can happen during (de)serialization.
type SerializationError int

const (

	// The format can not be identified
	SerializationUnsupportedFormatValue SerializationError = 0
	// The version of the data is not
	//   understood
	SerializationUnsupportedVersionValue SerializationError = 1
	// The given data may not exist in
	//   a proper serialization
	SerializationInvalidDataValue SerializationError = 2
)

// The categories of matrices relevant for GSK and GTK.
//
// Note that any category includes matrices of all later categories.
// So if you want to for example check if a matrix is a 2D matrix,
// `category &gt;= GSK_TRANSFORM_CATEGORY_2D` is the way to do this.
//
// Also keep in mind that rounding errors may cause matrices to not
// conform to their categories. Otherwise, matrix operations done via
// multiplication will not worsen categories. So for the matrix
// multiplication `C = A * B`, `category(C) = MIN (category(A), category(B))`.
type TransformCategory int

const (

	// The category of the matrix has not been
	//   determined.
	TransformCategoryUnknownValue TransformCategory = 0
	// Analyzing the matrix concluded that it does
	//   not fit in any other category.
	TransformCategoryAnyValue TransformCategory = 1
	// The matrix is a 3D matrix. This means that
	//   the w column (the last column) has the values (0, 0, 0, 1).
	TransformCategory3dValue TransformCategory = 2
	// The matrix is a 2D matrix. This is equivalent
	//   to graphene_matrix_is_2d() returning %TRUE. In particular, this
	//   means that Cairo can deal with the matrix.
	TransformCategory2dValue TransformCategory = 3
	// The matrix is a combination of 2D scale
	//   and 2D translation operations. In particular, this means that any
	//   rectangle can be transformed exactly using this matrix.
	TransformCategory2dAffineValue TransformCategory = 4
	// The matrix is a 2D translation.
	TransformCategory2dTranslateValue TransformCategory = 5
	// The matrix is the identity matrix.
	TransformCategoryIdentityValue TransformCategory = 6
)
