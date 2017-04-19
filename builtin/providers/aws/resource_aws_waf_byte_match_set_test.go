package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/acctest"
)

func TestAccAWSWafByteMatchSet_basic(t *testing.T) {
	var v waf.ByteMatchSet
	byteMatchSet := fmt.Sprintf("byteMatchSet-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafByteMatchSetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAWSWafByteMatchSetConfig(byteMatchSet),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafByteMatchSetExists("aws_waf_byte_match_set.byte_set", &v),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "name", byteMatchSet),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.#", "2"),
				),
			},
		},
	})
}

func TestAccAWSWafByteMatchSet_changeNameForceNew(t *testing.T) {
	var before, after waf.ByteMatchSet
	byteMatchSet := fmt.Sprintf("byteMatchSet-%s", acctest.RandString(5))
	byteMatchSetNewName := fmt.Sprintf("byteMatchSet-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafByteMatchSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafByteMatchSetConfig(byteMatchSet),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafByteMatchSetExists("aws_waf_byte_match_set.byte_set", &before),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "name", byteMatchSet),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.#", "2"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.text_transformation", "NONE"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.target_string", "badrefer1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.positional_constraint", "CONTAINS"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.2991901334.type", "HEADER"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.2991901334.data", "referer"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.text_transformation", "NONE"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.target_string", "badrefer2"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.positional_constraint", "CONTAINS"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.field_to_match.2991901334.type", "HEADER"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.field_to_match.2991901334.data", "referer"),
				),
			},
			{
				Config: testAccAWSWafByteMatchSetConfig(byteMatchSetNewName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafByteMatchSetExists("aws_waf_byte_match_set.byte_set", &after),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "name", byteMatchSetNewName),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.#", "2"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.text_transformation", "NONE"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.target_string", "badrefer1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.positional_constraint", "CONTAINS"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.2991901334.type", "HEADER"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.2991901334.data", "referer"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.text_transformation", "NONE"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.target_string", "badrefer2"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.positional_constraint", "CONTAINS"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.field_to_match.2991901334.type", "HEADER"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.field_to_match.2991901334.data", "referer"),
				),
			},
		},
	})
}

func TestAccAWSWafByteMatchSet_disappears(t *testing.T) {
	var v waf.ByteMatchSet
	byteMatchSet := fmt.Sprintf("byteMatchSet-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafByteMatchSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafByteMatchSetConfig(byteMatchSet),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafByteMatchSetExists("aws_waf_byte_match_set.byte_set", &v),
					testAccCheckAWSWafByteMatchSetDisappears(&v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAWSWafByteMatchSet_changeTuples(t *testing.T) {
	var before, after waf.ByteMatchSet
	name := fmt.Sprintf("byteMatchSet-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSWafByteMatchSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSWafByteMatchSetConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafByteMatchSetExists("aws_waf_byte_match_set.byte_set", &before),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "name", name),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.#", "2"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.text_transformation", "NONE"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.target_string", "badrefer1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.positional_constraint", "CONTAINS"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.2991901334.type", "HEADER"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.2991901334.data", "referer"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.text_transformation", "NONE"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.target_string", "badrefer2"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.positional_constraint", "CONTAINS"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.field_to_match.2991901334.type", "HEADER"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.839525137.field_to_match.2991901334.data", "referer"),
				),
			},
			{
				Config: testAccAWSWafByteMatchSetConfig_changedTuples(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSWafByteMatchSetExists("aws_waf_byte_match_set.byte_set", &after),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "name", name),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.#", "2"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.text_transformation", "NONE"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.target_string", "badrefer1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.positional_constraint", "CONTAINS"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.2991901334.type", "HEADER"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.2174619346.field_to_match.2991901334.data", "referer"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.1467312330.text_transformation", "HTML_ENTITY_DECODE"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.1467312330.target_string", "/images/daily-ad.jpg"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.1467312330.positional_constraint", "CONTAINS"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.1467312330.field_to_match.#", "1"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.1467312330.field_to_match.3756326843.type", "URI"),
					resource.TestCheckResourceAttr(
						"aws_waf_byte_match_set.byte_set", "byte_match_tuples.1467312330.field_to_match.3756326843.data", ""),
				),
			},
		},
	})
}

func testAccCheckAWSWafByteMatchSetDisappears(v *waf.ByteMatchSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*AWSClient).wafconn

		wr := newWafRetryer(conn, "global")
		_, err := wr.RetryWithToken(func(token *string) (interface{}, error) {
			req := &waf.UpdateByteMatchSetInput{
				ChangeToken:    token,
				ByteMatchSetId: v.ByteMatchSetId,
			}

			for _, ByteMatchTuple := range v.ByteMatchTuples {
				ByteMatchUpdate := &waf.ByteMatchSetUpdate{
					Action: aws.String("DELETE"),
					ByteMatchTuple: &waf.ByteMatchTuple{
						FieldToMatch:         ByteMatchTuple.FieldToMatch,
						PositionalConstraint: ByteMatchTuple.PositionalConstraint,
						TargetString:         ByteMatchTuple.TargetString,
						TextTransformation:   ByteMatchTuple.TextTransformation,
					},
				}
				req.Updates = append(req.Updates, ByteMatchUpdate)
			}

			return conn.UpdateByteMatchSet(req)
		})
		if err != nil {
			return errwrap.Wrapf("[ERROR] Error updating ByteMatchSet: {{err}}", err)
		}

		_, err = wr.RetryWithToken(func(token *string) (interface{}, error) {
			opts := &waf.DeleteByteMatchSetInput{
				ChangeToken:    token,
				ByteMatchSetId: v.ByteMatchSetId,
			}
			return conn.DeleteByteMatchSet(opts)
		})
		if err != nil {
			return errwrap.Wrapf("[ERROR] Error deleting ByteMatchSet: {{err}}", err)
		}

		return nil
	}
}

func testAccCheckAWSWafByteMatchSetExists(n string, v *waf.ByteMatchSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No WAF ByteMatchSet ID is set")
		}

		conn := testAccProvider.Meta().(*AWSClient).wafconn
		resp, err := conn.GetByteMatchSet(&waf.GetByteMatchSetInput{
			ByteMatchSetId: aws.String(rs.Primary.ID),
		})

		if err != nil {
			return err
		}

		if *resp.ByteMatchSet.ByteMatchSetId == rs.Primary.ID {
			*v = *resp.ByteMatchSet
			return nil
		}

		return fmt.Errorf("WAF ByteMatchSet (%s) not found", rs.Primary.ID)
	}
}

func testAccCheckAWSWafByteMatchSetDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_waf_byte_match_set" {
			continue
		}

		conn := testAccProvider.Meta().(*AWSClient).wafconn
		resp, err := conn.GetByteMatchSet(
			&waf.GetByteMatchSetInput{
				ByteMatchSetId: aws.String(rs.Primary.ID),
			})

		if err == nil {
			if *resp.ByteMatchSet.ByteMatchSetId == rs.Primary.ID {
				return fmt.Errorf("WAF ByteMatchSet %s still exists", rs.Primary.ID)
			}
		}

		// Return nil if the ByteMatchSet is already destroyed
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "WAFNonexistentItemException" {
				return nil
			}
		}

		return err
	}

	return nil
}

func testAccAWSWafByteMatchSetConfig(name string) string {
	return fmt.Sprintf(`
resource "aws_waf_byte_match_set" "byte_set" {
  name = "%s"
  byte_match_tuples {
    text_transformation = "NONE"
    target_string = "badrefer1"
    positional_constraint = "CONTAINS"
    field_to_match {
      type = "HEADER"
      data = "referer"
    }
  }

  byte_match_tuples {
    text_transformation = "NONE"
    target_string = "badrefer2"
    positional_constraint = "CONTAINS"
    field_to_match {
      type = "HEADER"
      data = "referer"
    }
  }
}`, name)
}

func testAccAWSWafByteMatchSetConfig_changedTuples(name string) string {
	return fmt.Sprintf(`
resource "aws_waf_byte_match_set" "byte_set" {
  name = "%s"
  byte_match_tuples {
    text_transformation = "NONE"
    target_string = "badrefer1"
    positional_constraint = "CONTAINS"
    field_to_match {
      type = "HEADER"
      data = "referer"
    }
  }

  byte_match_tuples {
    text_transformation = "HTML_ENTITY_DECODE"
    target_string = "/images/daily-ad.jpg"
    positional_constraint = "CONTAINS"
    field_to_match {
      type = "URI"
    }
  }
}`, name)
}
