package Instagram_media

import (
	"fmt"
	"math"
	"os"

	"github.com/aliforever/gista/media/geometry"

	mediaconstraints "github.com/aliforever/gista/media/media-constraints"

	"github.com/aliforever/gista/utils"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/media"
	"github.com/go-errors/errors"
)

type InstagramMedia struct {
	DefaultTmpPath                *string
	Debug                         bool
	InputFile                     string
	MinAspectRatio                *float64
	MaxAspectRatio                *float64
	AllowNewAspectDeviation       *bool
	HorCropFocus                  int
	VerCropFocus                  int
	BgColor                       []int
	Operation                     int
	TmpPath                       string
	Constraints                   media.Constraints
	OutputFile                    *string
	Details                       media.Media
	ForceTargetAspectRatio        *float64
	HasUserForceTargetAspectRatio bool
}

func NewInstagramMedia(inputFile string, options map[string]interface{}) (im *InstagramMedia, err error) {
	im = &InstagramMedia{}
	if !utils.FileOrFolderExists(inputFile) {
		err = errors.New(fmt.Sprintf(`Input file "%s" doesn"t exist.`, inputFile))
		return
	}
	targetFeed := constants.FeedTimeline
	if tf, ok := options["targetFeed"]; ok {
		if val, ok := tf.(string); !ok {
			err = errors.New(fmt.Sprintf("invalid option value type, %+v, expected value is string", tf))
			return
		} else {
			targetFeed = val
		}
	}
	if tf, ok := options["horCropFocus"]; ok {
		if val, ok := tf.(int); !ok || (val < -50 || val > 50) {
			err = errors.New("Horizontal crop focus must be between -50 and 50.")
			return
		} else {
			im.HorCropFocus = val
		}
	}
	if tf, ok := options["verCropFocus"]; ok {
		if val, ok := tf.(int); !ok || (val < -50 || val > 50) {
			err = errors.New("Vertical crop focus must be between -50 and 50.")
			return
		} else {
			im.VerCropFocus = val
		}
	}
	if tf, ok := options["operation"]; ok {
		if val, ok := tf.(int); !ok || (val != Crop && val != Expand) {
			err = errors.New(fmt.Sprintf("The operation must be one of the class constants CROP (%d) or EXPAND (%d).", Crop, Expand))
			return
		} else {
			im.Operation = val
		}
	}
	var MinAspectRatio *float64
	if tf, ok := options["MinAspectRatio"]; ok {
		if val, ok := tf.(float64); !ok {
			err = errors.New(fmt.Sprintf("invalid option value type, %+v, expected value is float64", tf))
			return
		} else {
			MinAspectRatio = &val
		}
	}
	var MaxAspectRatio *float64
	if tf, ok := options["MaxAspectRatio"]; ok {
		if val, ok := tf.(float64); !ok {
			err = errors.New(fmt.Sprintf("invalid option value type, %+v, expected value is float64", tf))
			return
		} else {
			MaxAspectRatio = &val
		}
	}
	var useRecommendedRatio *bool
	var forceAspectRatio *float64
	if tf, ok := options["forceAspectRatio"]; ok {
		if val, ok := tf.(float64); !ok {
			err = errors.New("Custom target aspect ratio must be a float64")
			return
		} else {
			fFalse := false
			forceAspectRatio = &val
			im.HasUserForceTargetAspectRatio = true
			useRecommendedRatio = &fFalse
		}
	}
	//var useRecommendedRatio *bool
	if tf, ok := options["useRecommendedRatio"]; ok {
		if val, ok := tf.(bool); !ok {
			err = errors.New(fmt.Sprintf("invalid option value type, %+v, expected value is bool", tf))
			return
		} else {
			useRecommendedRatio = &val
		}
	}
	var debug *bool
	if tf, ok := options["debug"]; ok {
		if val, ok := tf.(bool); !ok {
			err = errors.New(fmt.Sprintf("invalid option value type, %+v, expected value is bool", tf))
			return
		} else {
			debug = &val
		}
	}
	var allowNewAspectDeviation *bool
	if tf, ok := options["allowNewAspectDeviation"]; ok {
		if val, ok := tf.(bool); !ok {
			err = errors.New(fmt.Sprintf("invalid option value type, %+v, expected value is bool", tf))
			return
		} else {
			allowNewAspectDeviation = &val
		}
	}
	var bgColor []int
	if tf, ok := options["bgColor"]; ok {
		if val, ok := tf.([]int); !ok {
			err = errors.New(fmt.Sprintf("invalid option value type, %+v, expected value is []int", tf))
			return
		} else {
			bgColor = val
		}
	}
	var tmpPath *string
	if tf, ok := options["tmpPath"]; ok {
		if val, ok := tf.(string); !ok {
			err = errors.New(fmt.Sprintf("invalid option value type, %+v, expected value is string", tf))
			return
		} else {
			tmpPath = &val
		}
	}
	im.Debug = false
	if debug != nil && *debug {
		im.Debug = *debug
	}
	im.InputFile = inputFile
	im.Constraints = mediaconstraints.ConstraintsFactory{}.CreateFor(targetFeed)
	if !im.HasUserForceTargetAspectRatio && useRecommendedRatio == nil {
		if MinAspectRatio != nil || MaxAspectRatio != nil {
			fFalse := false
			useRecommendedRatio = &fFalse
		} else {
			b := im.Constraints.UseRecommendedRatioByDefault()
			useRecommendedRatio = &b
		}
	}

	if im.HasUserForceTargetAspectRatio && useRecommendedRatio != nil && *useRecommendedRatio {
		rec := im.Constraints.GetRecommendedRatio()
		im.ForceTargetAspectRatio = &rec
		deviation := im.Constraints.GetRecommendedRatioDeviation()
		min := *im.ForceTargetAspectRatio - deviation
		max := *im.ForceTargetAspectRatio - deviation
		MinAspectRatio = &min
		MaxAspectRatio = &max
	} else {
		/*
			im.forceTargetAspectRatio = userForceTargetAspectRatio;
			            allowedMinRatio = im.constraints.getMinAspectRatio();
			            allowedMaxRatio = im.constraints.getMaxAspectRatio();*/
		im.ForceTargetAspectRatio = forceAspectRatio
		allowedMinRatio := im.Constraints.GetMinAspectRatio()
		allowedMaxRatio := im.Constraints.GetMaxAspectRatio()

		if MinAspectRatio != nil && (*MinAspectRatio < allowedMinRatio || *MinAspectRatio > allowedMaxRatio) {
			err = errors.New(fmt.Sprintf(`Minimum aspect ratio must be between %.3f and %.3f.`, allowedMinRatio, allowedMaxRatio))
			return
		}

		if MinAspectRatio == nil {
			MinAspectRatio = &allowedMinRatio
		}

		if MaxAspectRatio != nil && (*MaxAspectRatio < allowedMinRatio || *MaxAspectRatio > allowedMaxRatio) {
			err = errors.New(fmt.Sprintf(`Maximum aspect ratio must be between %.3f and %.3f.`, allowedMinRatio, allowedMaxRatio))
			return
		}

		if MaxAspectRatio == nil {
			MaxAspectRatio = &allowedMaxRatio
		}

		if MinAspectRatio != nil && MaxAspectRatio != nil && *MinAspectRatio > *MaxAspectRatio {
			err = errors.New("Maximum aspect ratio must be greater than or equal to minimum.")
			return
		}
		if im.HasUserForceTargetAspectRatio {
			if MinAspectRatio != nil && *im.ForceTargetAspectRatio < *MinAspectRatio {
				err = errors.New(fmt.Sprintf(`Custom target aspect ratio (%.5f) must be greater than or equal to the minimum aspect ratio (%.5f).`, *im.ForceTargetAspectRatio, *MinAspectRatio))
			}
			if MaxAspectRatio != nil && *im.ForceTargetAspectRatio > *MaxAspectRatio {
				err = errors.New(fmt.Sprintf(`Custom target aspect ratio (%.5f) must be greater than or equal to the minimum aspect ratio (%.5f).`, *im.ForceTargetAspectRatio, *MaxAspectRatio))
			}
		}
	}
	im.MinAspectRatio = MinAspectRatio
	im.MaxAspectRatio = MaxAspectRatio
	im.AllowNewAspectDeviation = allowNewAspectDeviation

	if bgColor != nil && (len(bgColor) != 3) {
		err = errors.New("The background color must be a 3-element array [R, G, B].")
		return
	} else if bgColor == nil {
		bgColor = []int{255, 255, 255}
	}
	im.BgColor = bgColor

	if tmpPath == nil {
		tmpDir := os.TempDir()
		tmpPath = &tmpDir
		if im.DefaultTmpPath != nil {
			tmpPath = im.DefaultTmpPath
		}
	}

	if !utils.IsDirectory(*tmpPath) {
		err = errors.New(fmt.Sprintf(`Directory %s does not exist or is not writable.`, *tmpPath))
		return
	}
	im.TmpPath, err = utils.Realpath(*tmpPath)
	if err != nil {
		return
	}
	return
}

func (im *InstagramMedia) DeleteFile() (err error) {
	if im.OutputFile != nil && *im.OutputFile != im.InputFile && utils.FileOrFolderExists(*im.OutputFile) {
		err = os.Remove(*im.OutputFile)
		if err != nil {
			return
		}
		im.OutputFile = nil
	}
	return
}

func (im *InstagramMedia) GetFile() (path string, err error) {
	if im.OutputFile == nil {
		var shouldProcess bool
		shouldProcess, err = im.shouldProcess()
		if err != nil {
			return
		}
		im.OutputFile = &im.InputFile
		if shouldProcess {
			im.process()
		}
	}
	return
}

func (im *InstagramMedia) shouldProcess() (result bool, err error) {
	inputAspectRatio := im.Details.GetAspectRatio()

	if im.MinAspectRatio != nil && inputAspectRatio < *im.MinAspectRatio {
		result = true
		return
	}

	if im.MaxAspectRatio != nil && inputAspectRatio > *im.MaxAspectRatio {
		result = true
		return
	}

	if im.HasUserForceTargetAspectRatio {
		if *im.ForceTargetAspectRatio == 1.0 {
			if inputAspectRatio != 1.0 {
				result = true
				return
			}
		} else {
			acceptableDeviation := 0.003
			acceptableMinAspectRatio := *im.ForceTargetAspectRatio - acceptableDeviation
			acceptableMaxAspectRatio := *im.ForceTargetAspectRatio + acceptableDeviation
			if inputAspectRatio < acceptableMinAspectRatio || inputAspectRatio > acceptableMaxAspectRatio {
				result = true
				return
			}
		}
	}
	err = im.Details.Validate(im.Constraints)
	return
}

func (im *InstagramMedia) process() (path string, err error) {
	inputCanvas := geometry.NewDimensions(im.Details.GetWidth(), im.Details.GetHeight())
	var canvasInfo map[string]interface{}
	canvasInfo, err = im.calculateNewCanvas(
		im.Operation, inputCanvas.GetWidth(), inputCanvas.GetHeight(), im.isMod2CanvasRequired(), im.Details.GetMinAllowedWidth(), im.Details.GetMaxAllowedWidth(),
		im.MinAspectRatio, im.MaxAspectRatio, im.ForceTargetAspectRatio, im.AllowNewAspectDeviation)
	if err != nil {
		return
	}
	outputCanvas := canvasInfo["canvas"].(geometry.Dimensions)
	if im.Operation == Crop {
		idealCanvas := geometry.NewDimensions(outputCanvas.GetWidth()-canvasInfo["mod2WidthDiff"].(int), outputCanvas.GetHeight()-canvasInfo["mod2HeightDiff"].(int))
		idealWidthScale := float64(idealCanvas.GetWidth() / inputCanvas.GetWidth())
		idealHeightScale := float64(idealCanvas.GetHeight() / inputCanvas.GetHeight())
		text := "CROP: Analyzing Original Input Canvas Size"
		im.debugDimensions(inputCanvas.GetWidth(), inputCanvas.GetHeight(), &text)
		text = "CROP: Analyzing Ideally Cropped (Non-Mod2-adjusted) Output Canvas Size"
		im.debugDimensions(idealCanvas.GetWidth(), idealCanvas.GetHeight(), &text)
		text = "CROP: Scale of Ideally Cropped Canvas vs Input Canvas"
		im.debugText(text, "width=%.8f, height=%.8f", idealWidthScale, idealHeightScale)
	}
	// TODO:
	return
}

func (im *InstagramMedia) isMod2CanvasRequired() bool {
	return false
}

func (im *InstagramMedia) calculateNewCanvas(operation, inputWidth, inputHeight int, isMod2CanvasRequired bool, minWidth, maxWidth int, minAspectRatio, maxAspectRatio, forceTargetAspectRatio *float64, allowNewAspectDeviation *bool) (result map[string]interface{}, err error) {
	if forceTargetAspectRatio != nil {
		im.debugText("SPECIAL_PARAMETERS: Forced Target Aspect Ratio", "forceTargetAspectRatio=%.5f", forceTargetAspectRatio)
	}
	targetWidth := inputWidth
	targetHeight := inputHeight
	targetAspectRatio := float64(inputWidth) / float64(inputHeight)
	text := "CANVAS_INPUT: Input Canvas Size"
	im.debugDimensions(targetWidth, targetHeight, &text)
	if (minAspectRatio != nil && targetAspectRatio < *minAspectRatio) || (forceTargetAspectRatio != nil && targetAspectRatio < *forceTargetAspectRatio) {
		targetAspectRatio = *minAspectRatio
		if forceTargetAspectRatio != nil {
			targetAspectRatio = *forceTargetAspectRatio
		}

		if operation == Crop {
			targetHeight = int(math.Floor(float64(targetWidth) / targetAspectRatio))
			text := "Applying Forced Aspect for INPUT < TARGET"
			if forceTargetAspectRatio == nil {
				text = "Aspect Was < MIN"
			}
			text = fmt.Sprintf("CANVAS_CROPPED: %s", text)
			im.debugDimensions(targetWidth, targetHeight, &text)
		} else if operation == Expand {
			targetWidth = int(math.Ceil(float64(targetHeight) * targetAspectRatio))
			text := "Applying Forced Aspect for INPUT < TARGET"
			if forceTargetAspectRatio == nil {
				text = "Aspect Was < MIN"
			}
			text = fmt.Sprintf("CANVAS_EXPANDED: %s", text)
			im.debugDimensions(targetWidth, targetHeight, &text)
		}
	} else if (maxAspectRatio != nil && targetAspectRatio > *maxAspectRatio) || (forceTargetAspectRatio != nil && targetAspectRatio > *forceTargetAspectRatio) {
		targetAspectRatio = *maxAspectRatio
		if forceTargetAspectRatio != nil {
			targetAspectRatio = *forceTargetAspectRatio
		}
		if operation == Crop {
			targetWidth = int(math.Floor(float64(targetHeight) * targetAspectRatio))
			text := "Applying Forced Aspect for INPUT > TARGET"
			if forceTargetAspectRatio != nil {
				text = "Aspect Was > MAX"
			}
			text = fmt.Sprintf("CANVAS_CROPPED: %s", text)
			im.debugDimensions(targetWidth, targetHeight, &text)
		} else if operation == Expand {
			targetHeight = int(math.Ceil(float64(targetWidth) / targetAspectRatio))
			text := "Applying Forced Aspect for INPUT > TARGET"
			if forceTargetAspectRatio != nil {
				text = "Aspect Was > MAX"
			}
			text = fmt.Sprintf("CANVAS_EXPANDED: %s", text)
			im.debugDimensions(targetWidth, targetHeight, &text)
		}
	} else {
		text := "CANVAS: Aspect Ratio Already Legal"
		im.debugDimensions(targetWidth, targetHeight, &text)
	}
	val := 0 - targetAspectRatio
	if minAspectRatio != nil {
		val = *minAspectRatio - targetAspectRatio
	}
	minAspectDistance := math.Abs(float64(val))
	val = 9999999 - targetAspectRatio
	if maxAspectRatio != nil {
		val = *maxAspectRatio - targetAspectRatio
	}
	maxAspectDistance := math.Abs(float64(9999999))
	isClosestToMinAspect := minAspectDistance <= maxAspectDistance
	useFloorHeightRecalc := isClosestToMinAspect
	if targetAspectRatio == 1.0 && targetWidth != targetHeight {
		val := math.Max(float64(targetWidth), float64(targetHeight))
		if operation == Crop {
			val = math.Min(float64(targetWidth), float64(targetHeight))
		}
		targetWidth, targetHeight = int(val), int(val)
		text := "CANVAS_SQUARIFY: Fixed Badly Generated Square"
		im.debugDimensions(targetWidth, targetHeight, &text)
	}
	if targetWidth > maxWidth {
		targetWidth = maxWidth
		text := "CANVAS_WIDTH: Width Was > MAX"
		im.debugDimensions(targetWidth, targetHeight, &text)
		targetHeight = im.accurateHeightRecalc(useFloorHeightRecalc, targetAspectRatio, targetWidth)
		text = "CANVAS_WIDTH: Height Recalc From Width & Aspect"
		im.debugDimensions(targetWidth, targetHeight, &text)
	} else if targetWidth < minWidth {
		targetWidth = minWidth
		text = "CANVAS_WIDTH: Width Was < MIN"
		im.debugDimensions(targetWidth, targetHeight, &text)
		targetHeight = im.accurateHeightRecalc(useFloorHeightRecalc, targetAspectRatio, targetWidth)
		text = "CANVAS_WIDTH: Height Recalc From Width & Aspect"
		im.debugDimensions(targetWidth, targetHeight, &text)
	}
	mod2WidthDiff, mod2HeightDiff := 0, 0
	if isMod2CanvasRequired && (!im.isNumberMod2(targetWidth) || !im.isNumberMod2(targetHeight)) {
		// Calculate the Mod2-adjusted final canvas size.
		var mod2Canvas geometry.Dimensions
		mod2Canvas, err = im.calculateAdjustedMod2Canvas(
			inputWidth,
			inputHeight,
			useFloorHeightRecalc,
			targetWidth,
			targetHeight,
			targetAspectRatio,
			minWidth,
			maxWidth,
			minAspectRatio,
			maxAspectRatio,
			*allowNewAspectDeviation,
		)
		if err != nil {
			return
		}
		mod2WidthDiff = mod2Canvas.GetWidth() - targetWidth
		mod2HeightDiff = mod2Canvas.GetHeight() - targetHeight
		im.debugText("CANVAS: Mod2 Difference Stats", "width=%s, height=%s", mod2WidthDiff, mod2HeightDiff)

		targetWidth = mod2Canvas.GetWidth()
		targetHeight = mod2Canvas.GetHeight()
		text = "CANVAS: Updated From Mod2 Result"
		im.debugDimensions(targetWidth, targetHeight, &text)
	}
	canvas := geometry.NewDimensions(targetWidth, targetHeight)
	text = "CANVAS_OUTPUT: Final Output Canvas Size"
	im.debugDimensions(targetWidth, targetHeight, &text)
	isIllegalRatio := (minAspectRatio != nil && canvas.GetAspectRatio() < *minAspectRatio) || (maxAspectRatio != nil && canvas.GetAspectRatio() > *maxAspectRatio)
	if canvas.GetWidth() < 1 || canvas.GetHeight() < 1 {
		err = errors.New(fmt.Sprintf(`Canvas calculation failed. Target width (%d) or height (%d) less than one pixel.`, canvas.GetWidth(), canvas.GetHeight()))
		return
	} else if canvas.GetWidth() < minWidth {
		err = errors.New(fmt.Sprintf(`Canvas calculation failed. Target width (%d) less than minimum allowed (%d).`, canvas.GetWidth(), minWidth))
		return
	} else if canvas.GetWidth() > maxWidth {
		err = errors.New(fmt.Sprintf(`Canvas calculation failed. Target width (%d) greater than maximum allowed (%d).`, canvas.GetWidth(), maxWidth))
		return
	} else if isIllegalRatio {
		if !*allowNewAspectDeviation {
			valMin := 0.0
			if minAspectRatio != nil {
				val = *minAspectRatio
			}
			valMax := math.Inf(1)
			if maxAspectRatio != nil {
				valMax = *maxAspectRatio
			}
			err = errors.New(fmt.Sprintf(`Canvas calculation failed. Unable to reach target aspect ratio range during output canvas generation. The range of allowed aspect ratios is too narrow (%.8f - %.8f). We achieved a ratio of %.8f.`, valMin, valMax, canvas.GetAspectRatio()))
			return
		} else {
			// The user wants us to allow "near-misses", so we proceed...
			text = "CANVAS_FINAL: Allowing Deviating Aspect Ratio"
			im.debugDimensions(canvas.GetWidth(), canvas.GetHeight(), &text)
		}
	}
	result = map[string]interface{}{
		"canvas":         canvas,
		"mod2WidthDiff":  mod2WidthDiff,
		"mod2HeightDiff": mod2HeightDiff,
	}
	return
}

func (im *InstagramMedia) calculateAdjustedMod2Canvas(inputWidth, inputHeight int, useFloorHeightRecalc bool, targetWidth, targetHeight int, targetAspectRatio float64, minWidth, maxWidth int, minAspectRatio, maxAspectRatio *float64, allowNewAspectDeviation bool) (d geometry.Dimensions, err error) {
	mod2Width := targetWidth
	mod2Height := targetHeight
	text := "MOD2_CANVAS: Current Canvas Size"
	im.debugDimensions(mod2Width, mod2Height, &text)
	canCutWidth := mod2Width > minWidth
	if !im.isNumberMod2(mod2Width) {
		plus := 1
		if canCutWidth {
			plus = -1
		}
		mod2Width += plus
		text = "MOD2_CANVAS: Width Mod2Fix"
		im.debugDimensions(mod2Width, mod2Height, &text)

		mod2Height = im.accurateHeightRecalc(useFloorHeightRecalc, targetAspectRatio, mod2Width)
		text = "MOD2_CANVAS: Height Recalc From Width & Aspect"
		im.debugDimensions(mod2Width, mod2Height, &text)
	}
	if !im.isNumberMod2(mod2Height) {
		plus := 1
		if canCutWidth {
			plus = -1
		}
		mod2Height += plus
		text = "MOD2_CANVAS: Height Mod2Fix"
		im.debugDimensions(mod2Width, mod2Height, &text)
	}
	heightAlternatives := map[string][]map[string]interface{}{
		"perfect": {},
		"stretch": {},
		"bad":     {},
	}
	offsetPriorities := []int{0, 2, -2, 4, -4, 6, -6}
	for _, offset := range offsetPriorities {
		offsetMod2Height := mod2Height + offset
		offsetMod2AspectRatio := float64(mod2Width / offsetMod2Height)

		isLegalRatio := (minAspectRatio == nil || offsetMod2AspectRatio >= *minAspectRatio) && (maxAspectRatio == nil || offsetMod2AspectRatio <= *maxAspectRatio)

		stretchAmount := math.Max(0, float64(offsetMod2Height-inputHeight))

		ratioDeviation := math.Abs(float64(offsetMod2AspectRatio - targetAspectRatio))

		rating := "bad"
		if isLegalRatio && stretchAmount == 0 {
			rating = "perfect"
		} else if isLegalRatio {
			rating = "stretch"
		}
		heightAlternatives[rating] = []map[string]interface{}{
			{
				"offset":         offset,
				"height":         offsetMod2Height,
				"ratio":          offsetMod2AspectRatio,
				"isLegalRatio":   isLegalRatio,
				"stretchAmount":  stretchAmount,
				"ratioDeviation": ratioDeviation,
				"rating":         rating,
			},
		}
		p := `""`
		if offset > 0 {
			p = `"+"`
		}
		text := fmt.Sprintf(`MOD2_CANVAS_CHECK: Testing Height Mod2Ratio (h%s%d = %s)`, p, offset, rating)
		im.debugDimensions(mod2Width, offsetMod2Height, &text)
	}
	var bestHeight map[string]interface{}
	ratings := []string{"perfect", "stretch", "bad"}
	for _, rating := range ratings {
		if val, ok := heightAlternatives[rating]; ok && val != nil {
			var least float64
			for _, v := range val {
				if least == 0 {
					least = v["ratioDeviation"].(float64)
					bestHeight = v
				} else {
					if v["ratioDeviation"].(float64) < least {
						least = v["ratioDeviation"].(float64)
						bestHeight = v
					}
				}
			}
			break
		}
	}
	mod2Height = bestHeight["height"].(int)
	val := `""`
	if bestHeight["offset"].(int) >= 0 {
		val = `"+"`
	}
	text = fmt.Sprintf(`MOD2_CANVAS: Selected Most Ideal Height Mod2Ratio (h%s%d = %s)`, val, bestHeight["offset"].(int), bestHeight["rating"].(string))
	im.debugDimensions(mod2Width, mod2Height, &text)

	if bestHeight["rating"].(string) == "bad" {
		if !allowNewAspectDeviation {
			minVal := 0.0
			if minAspectRatio != nil {
				minVal = *minAspectRatio
			}
			maxVal := math.Inf(1)
			if maxAspectRatio != nil {
				maxVal = *maxAspectRatio
			}
			text = fmt.Sprintf("Canvas calculation failed. Unable to reach target aspect ratio range during Mod2 canvas conversion. The range of allowed aspect ratios is too narrow (%.8f - %.8f). We achieved a ratio of %.8f.", minVal, maxVal, float64(mod2Width/mod2Height))
			err = errors.New(text)
			return
		} else {
			val := `""`
			if bestHeight["offset"].(int) >= 0 {
				val = `"+"`
			}
			text := fmt.Sprintf(`MOD2_CANVAS: Allowing Deviating Height Mod2Ratio (h%s%d = %s)`, val, bestHeight["offset"].(int), bestHeight["rating"].(string))
			im.debugDimensions(mod2Width, mod2Height, &text)
		}
	}
	d = geometry.NewDimensions(mod2Width, mod2Height)
	return
}

func (im *InstagramMedia) isNumberMod2(number int) bool {
	return number%2 == 0
}

func (im *InstagramMedia) accurateHeightRecalc(useFloorHeightRecalc bool, targetAspectRatio float64, targetWidth int) (targetHeight int) {
	targetHeight = int(math.Ceil(float64(targetWidth) / targetAspectRatio))
	if useFloorHeightRecalc {
		targetHeight = int(math.Floor(float64(targetWidth) / targetAspectRatio))
	}
	return
}

func (im *InstagramMedia) debugText(stepDescription, formatMessage string, args ...interface{}) {
	if !im.Debug {
		return
	}
	fmt.Println(fmt.Sprintf(`[[1;33m%s[0m] `+formatMessage+`\n`, stepDescription, args))
}

func (im *InstagramMedia) debugDimensions(width, height int, stepDescription *string) {
	if !im.Debug {
		return
	}
	description := "DEBUG"
	if stepDescription != nil {
		description = *stepDescription
	}
	fmt.Println(fmt.Sprintf(`[[1;33m%s[0m] w=%s h=%s (aspect %.8f)\n`, description, width, height, float64(width/height)))
}
