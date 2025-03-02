package config

import (
	"github.com/RyseUp/in-game-item-system/internal/models"
	"log"
)

func SeedDatabase() {
	items := []models.Item{
		{ItemID: "item_001", Name: "Magic Sword", Description: "A legendary sword with magical power."},
		{ItemID: "item_002", Name: "Healing Potion", Description: "Restores health over time."},
		{ItemID: "item_003", Name: "Dragon Shield", Description: "Protects against dragon fire."},
		{ItemID: "item_004", Name: "Silver Bow", Description: "A bow crafted with enchanted silver."},
		{ItemID: "item_005", Name: "Mana Elixir", Description: "Restores magical energy instantly."},
		{ItemID: "item_006", Name: "Phoenix Feather", Description: "Used to revive fallen warriors."},
		{ItemID: "item_007", Name: "Thunder Axe", Description: "An axe that crackles with electricity."},
		{ItemID: "item_008", Name: "Invisibility Cloak", Description: "Grants temporary invisibility."},
		{ItemID: "item_009", Name: "Knight’s Helmet", Description: "A sturdy helmet worn by royal knights."},
		{ItemID: "item_010", Name: "Teleportation Scroll", Description: "Teleports the user to a known location."},
		{ItemID: "item_011", Name: "Ring of Strength", Description: "Increases physical attack power."},
		{ItemID: "item_012", Name: "Mystic Orb", Description: "Enhances magical abilities."},
		{ItemID: "item_013", Name: "Ice Spear", Description: "A weapon infused with ice magic."},
		{ItemID: "item_014", Name: "Fireball Staff", Description: "Casts powerful fireball spells."},
		{ItemID: "item_015", Name: "Earthquake Hammer", Description: "Causes ground tremors upon impact."},
		{ItemID: "item_016", Name: "Wind Dagger", Description: "A lightweight dagger with wind enchantment."},
		{ItemID: "item_017", Name: "Shadow Boots", Description: "Enhances speed and stealth."},
		{ItemID: "item_018", Name: "Tome of Knowledge", Description: "Contains ancient spells and wisdom."},
		{ItemID: "item_019", Name: "Golden Armor", Description: "A heavy armor forged with enchanted gold."},
		{ItemID: "item_020", Name: "Summoner’s Bell", Description: "Calls forth a mystical beast for battle."},
	}

	for _, item := range items {
		var existing models.Item
		result := DB.Where("item_id = ?", item.ItemID).First(&existing)

		if result.Error != nil {
			if err := DB.Create(&item).Error; err != nil {
				log.Printf("Failed to insert item %s: %v", item.ItemID, err)
			} else {
				log.Printf("Inserted item: %s", item.Name)
			}
		}
	}

	inventories := []models.Inventory{
		{InventoryID: "inv_001", UserID: "user_001", ItemID: "item_001", Quantity: 10},
		{InventoryID: "inv_002", UserID: "user_001", ItemID: "item_002", Quantity: 5},
	}

	for _, inv := range inventories {
		var existing models.Inventory
		result := DB.Where("user_id = ? AND item_id = ?", inv.UserID, inv.ItemID).First(&existing)
		if result.Error != nil {
			if err := DB.Create(&inv).Error; err != nil {
				log.Printf("Failed to insert inventory for user %s and item %s: %v", inv.UserID, inv.ItemID, err)
			} else {
				log.Printf("Inserted inventory for user %s and item %s", inv.UserID, inv.ItemID)
			}
		}
	}

}
