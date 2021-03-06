package json

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/tayusa/notugly_backend/internal/domain"
)

type coordinateRepository struct {
}

func (c *coordinateRepository) FindById(
	_ context.Context, id int) (domain.Coordinate, error) {

	coordinates, err := GetCoordinates(GET)
	if err != nil {
		return domain.Coordinate{}, err
	}

	for _, coordinate := range coordinates {
		if coordinate.Id == id {
			return coordinate, nil
		}
	}

	return domain.Coordinate{}, fmt.Errorf("err no rows")
}

func (c *coordinateRepository) GetAtRandom(_ context.Context) (domain.Coordinate, error) {
	coordinates, err := GetCoordinates(GET)
	if err != nil {
		return domain.Coordinate{}, err
	}

	if len(coordinates) > 0 {
		return coordinates[0], nil
	}

	return domain.Coordinate{}, fmt.Errorf("err no rows")
}

func (c *coordinateRepository) FindFavoriteByCoordinateIdAndUserId(
	_ context.Context, id int, uid string) (domain.Favorite, error) {

	favorites, err := GetFavorites(GET)
	if err != nil {
		return domain.Favorite{}, err
	}

	for _, favorite := range favorites {
		if favorite.CoordinateId == id && favorite.UserId == uid {
			return favorite, nil
		}
	}

	return domain.Favorite{}, fmt.Errorf("err no rows")
}

func (c *coordinateRepository) FindByUserId(
	_ context.Context, userId string) ([]domain.Coordinate, error) {

	coordinates, err := GetCoordinates(GET)
	if err != nil {
		return []domain.Coordinate{}, err
	}

	output := make([]domain.Coordinate, 0)
	for _, coordinate := range coordinates {
		if coordinate.UserId == userId {
			output = append(output, coordinate)
		}
	}
	return output, nil
}

func (c *coordinateRepository) Store(_ context.Context, _ domain.Coordinate) (int64, error) {
	return 0, nil
}

func (c *coordinateRepository) Delete(_ context.Context, _ int) error {
	return nil
}

func NewCoordinateRepository() *coordinateRepository {
	return &coordinateRepository{}
}

func GetCoordinates(method int) ([]domain.Coordinate, error) {
	bytes, err := ioutil.ReadFile(
		filepath.Join(jsonPath, "coordinates", methods[method]+".json"))
	if err != nil {
		return []domain.Coordinate{}, err
	}

	var coordinates []domain.Coordinate
	if err = json.Unmarshal(bytes, &coordinates); err != nil {
		return []domain.Coordinate{}, err
	}
	return coordinates, nil
}
