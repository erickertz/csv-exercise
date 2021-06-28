resource "google_storage_bucket_iam_policy" "tfer--scior-002D-terraform" {
  bucket = "b/scior-terraform"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tfer--scoir-002D-csv" {
  bucket = "b/scoir-csv"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tfer--scoir-002D-json" {
  bucket = "b/scoir-json"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}
