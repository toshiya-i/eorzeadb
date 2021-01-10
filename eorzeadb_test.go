package eorzeadb

import (
	"reflect"
	"testing"
)

func TestNewItemSuccessCase1(t *testing.T) {
	// エクサーク・ストライカートップス
	url := "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/038fab75d02/"

	result, err := NewItem(url)
	if err != nil {
		t.Fatalf("failed err %#v", err)
	}
	if result.URL != url {
		t.Fatalf("failed result %#v", result)
	}
	if result.Name != "エクサーク・ストライカートップス" {
		t.Fatalf("failed result %#v", result)
	}
	if result.Category != "胴防具" {
		t.Fatalf("failed result %#v", result)
	}
	if result.EquipmentJob != "格闘士 モンク 侍" {
		t.Fatalf("failed result %#v", result)
	}
	if result.RecipeURL != "https://jp.finalfantasyxiv.com/lodestone/playguide/db/recipe/9988bef02b8/" {
		t.Fatalf("failed result %#v", result)
	}
	t.Logf("result %#v", result)
}

func TestNewItemSuccessCase2(t *testing.T) {
	// ジュラルミンインゴット
	url := "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/52fd8b325e4/"

	result, err := NewItem(url)
	if err != nil {
		t.Fatalf("failed err %#v", err)
	}
	if result.URL != url {
		t.Fatalf("failed result %#v", result)
	}
	if result.Name != "ジュラルミンインゴット" {
		t.Fatalf("failed result %#v", result)
	}
	if result.Category != "金属材" {
		t.Fatalf("failed result %#v", result)
	}
	if result.EquipmentJob != "" {
		t.Fatalf("failed result %#v", result)
	}
	if result.RecipeURL != "https://jp.finalfantasyxiv.com/lodestone/playguide/db/recipe/9d433def648/" {
		t.Fatalf("failed result %#v", result)
	}
	t.Logf("result %#v", result)
}

func TestNewRecipe(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *Recipe
		wantErr bool
	}{
		{
			name: "case1_success_エクサーク・ストライカートップス",
			args: args{
				url: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/recipe/9988bef02b8/",
			},
			want: &Recipe{
				URL:          "https://jp.finalfantasyxiv.com/lodestone/playguide/db/recipe/9988bef02b8/",
				TotalCrafted: "1",
				Materials: [5]Material{
					{
						URL:  "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/8f264762e2a/",
						Name: "クチナシ染め布",
						Num:  "2",
					},
					{
						URL:  "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/52fd8b325e4/",
						Name: "ジュラルミンインゴット",
						Num:  "2",
					},
					{
						URL:  "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/55a11a47140/",
						Name: "剛力の幻水G4",
						Num:  "2",
					},
					{
						URL:  "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/64007c6b68b/",
						Name: "シースワローレザー",
						Num:  "1",
					},
					{
						URL:  "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/b7b850d3bf5/",
						Name: "雷鳴の霊砂",
						Num:  "2",
					},
				},
				ItemURL: "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/038fab75d02/",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRecipe(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRecipe() = %v, want %v", got, tt.want)
			}
			t.Logf("NewRecipe() = %v, want %v", got, tt.want)
		})
	}
}
